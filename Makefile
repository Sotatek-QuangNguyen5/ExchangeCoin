solc:
	solc --abi --bin --overwrite smartcontract/ExchangeCoin.sol -o build

abigen:
	abigen --bin=build/VerifySignature.bin --abi=build/VerifySignature.abi --pkg=exchangecoin --out=exchangecoin/ExchangeCoin.go

deploy:
	go build

run:
	./servercoin