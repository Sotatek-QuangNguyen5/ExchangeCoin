package readfile

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