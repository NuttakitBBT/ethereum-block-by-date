package getDate

import (
	"fmt"
	"context"
    "log"
	"math/big"
	"github.com/ethereum/go-ethereum/ethclient"
)



func getDate(date string) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/8e832100bb9d451887b920a0935dc120")
    if err != nil {
        log.Fatal(err)
    }

    header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
	curBlock := header.Number.Uint64()
	blockNumber := big.NewInt(curBlock)
	unixCB := curBlock.Time()
}