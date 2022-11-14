package contract

import (
    
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
    "fmt"
	"servercoin/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	Client *ethclient.Client
	Auth   *bind.TransactOpts
    fromAddress common.Address
    privateKeyLocal *ecdsa.PrivateKey
)

func Init() {

    urlrpc := utils.ReadUrlRpc(1)
    client, err := ethclient.Dial(urlrpc)
    if err != nil {
        log.Fatal(err)
    }
    Client = client

	privateKeyString := utils.ReadPrivateKey(1)

    privateKey, err := crypto.HexToECDSA(privateKeyString)
    if err != nil {
        log.Fatal(err)
    }
    privateKeyLocal = privateKey

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := Client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

	chainID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
		log.Fatal(err)
	}
    Auth = auth
    Auth.Nonce = big.NewInt(int64(nonce))
    Auth.Value = big.NewInt(int64(0))
    Auth.GasLimit = uint64(3000000)
    Auth.GasPrice = gasPrice
}

func SetValue(value *big.Int) {

    Auth.Value = value
}

func SetNonce() {

    nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }
    Auth.Nonce = big.NewInt(int64(nonce))
}

func GenerateSignature(addr string, mess string, am string, one int64) string {

    message := ([]byte)(mess)
    amount := ([]byte)(am)
    address := ([]byte)(addr)
    nonce := big.NewInt(one).Bytes()
    hash, ethHash := PacketWithEth(address, message, amount, nonce)
    fmt.Println("Msg Hash : ", hash.Hex())
    fmt.Println("ETH Hash : ", ethHash.Hex())

    signature, err := crypto.Sign(ethHash.Bytes(), privateKeyLocal)
    if err != nil {

        log.Fatal(err)
    }
    
    if signature[64] == 0 || signature[64] == 1 {

        signature[64] += 27
    }
    
    // fmt.Println(signature)
    // fmt.Println("Signature : ", hexutil.Encode(signature))
    return hexutil.Encode(signature)
}

func PacketWithEth(address []byte, message []byte, amount []byte, nonce []byte) (common.Hash, common.Hash) {

    hash := crypto.Keccak256Hash(address, message, amount, nonce)

    return hash, crypto.Keccak256Hash(
        []byte("\x19Ethereum Signed Message:\n32"),
        hash.Bytes(),
    )
}

func VerifySignature(ethHash common.Hash, signature []byte) common.Address {

    if signature[64] == 27 || signature[64] == 28 {

        signature[64] -= 27
    }
    rpk, err := crypto.SigToPub(ethHash.Bytes(), signature)
    if err != nil {
        
        log.Fatal("Err : ", err)
    }
    pubKey := crypto.PubkeyToAddress(*rpk)
    return pubKey
}
