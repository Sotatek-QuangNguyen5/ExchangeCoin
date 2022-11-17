package utils

import (
	
    "bufio"
    "log"
    "os"
)

func ReadContract(number int) (string) {


    f, err := os.Open("./text/contract.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var contractAddress string
    cnt := 1

    for scanner.Scan() {

		contractAddress = scanner.Text()
        if cnt == number {

            break
        }
		cnt += 1
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return contractAddress
}

func ReadPrivateKey(number int) (string) {


    f, err := os.Open("./text/privatekey.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var privateKey string
    cnt := 1

    for scanner.Scan() {

		privateKey = scanner.Text()
        if cnt == number {

            break
        }
        cnt += 1
    }

    if err := scanner.Err(); err != nil {

        log.Fatal(err)
    }

	return privateKey
}

func ReadUrlRpc(number int) (string) {


    f, err := os.Open("./text/urlrpc.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var privateKey string
    cnt := 1

    for scanner.Scan() {

		privateKey = scanner.Text()
        if cnt == number {

            break
        }
        cnt += 1
    }

    if err := scanner.Err(); err != nil {

        log.Fatal(err)
    }

	return privateKey
}

func ReadPublicKey(number int) (string) {


    f, err := os.Open("./text/publickey.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var publicKey string
    cnt := 1

    for scanner.Scan() {

		publicKey = scanner.Text()
        if cnt == number {

            break
        }
        cnt += 1
    }

    if err := scanner.Err(); err != nil {

        log.Fatal(err)
    }

	return publicKey
}

func ReadSignature(number int) (string) {


    f, err := os.Open("./text/signature.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var signature string
    cnt := 1

    for scanner.Scan() {

		signature = scanner.Text()
        if cnt == number {

            break
        }
        cnt += 1
    }

    if err := scanner.Err(); err != nil {

        log.Fatal(err)
    }

	return signature
}

func ReadETHHash(number int) (string) {


    f, err := os.Open("./text/ethhash.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var ethHash string
    cnt := 1

    for scanner.Scan() {

		ethHash = scanner.Text()
        if cnt == number {

            break
        }
        cnt += 1
    }

    if err := scanner.Err(); err != nil {

        log.Fatal(err)
    }

	return ethHash
}

func ReadAmount(number int) (string) {


    f, err := os.Open("./text/amount.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var amount string
    cnt := 1

    for scanner.Scan() {

		amount = scanner.Text()
        if cnt == number {

            break
        }
        cnt += 1
    }

    if err := scanner.Err(); err != nil {

        log.Fatal(err)
    }

	return amount
}

func ReadMessage(number int) (string) {


    f, err := os.Open("./text/message.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var message string
    cnt := 1

    for scanner.Scan() {

		message = scanner.Text()
        if cnt == number {

            break
        }
        cnt += 1
    }

    if err := scanner.Err(); err != nil {

        log.Fatal(err)
    }

	return message
}