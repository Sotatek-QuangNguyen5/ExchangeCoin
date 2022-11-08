package initstate

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
    "fmt"
	"servercoin/readfile"

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

    client, err := ethclient.Dial("HTTP://127.0.0.1:8545")
    if err != nil {
        log.Fatal(err)
    }
    Client = client

	privateKeyString := readfile.ReadPrivateKey()

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

func SetValue(value int64) {

    Auth.Value = big.NewInt(value)
}

func SetNonce() {

    nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }
    Auth.Nonce = big.NewInt(int64(nonce))
}

func GenerateSignature() {

    // to := common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")
    // amount := *big.NewInt(1e4)
    message := ([]byte)("oke")
    // nonce := *big.NewInt(119)
    // hash := crypto.Keccak256Hash(to.Bytes(), amount.Bytes(), message, nonce.Bytes())
    hash := crypto.Keccak256Hash(message)
    fmt.Println("HashHex : ", hash.Hex())

    signature, err := crypto.Sign(hash.Bytes(), privateKeyLocal)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Signature : ", hexutil.Encode(signature))
}