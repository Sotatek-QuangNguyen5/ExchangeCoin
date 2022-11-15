solc:
	solc --abi --bin --overwrite smartcontract/ExchangeCoin.sol -o build

abigen:
	abigen --bin=build/ExchangeCoin.bin --abi=build/ExchangeCoin.abi --pkg=exchangecoin --out=exchangecoin/ExchangeCoin.go

deploy:
	go build

run:
	./servercoin

test:
	go run main.go