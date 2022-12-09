package main

import (
    "context"
    "fmt"
    "log"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/8e832100bb9d451887b920a0935dc120")
    if err != nil {
        log.Fatal(err)
    }

    header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
	getDate.SayHi()
    fmt.Println(header.Number.String()) // 5671744
}