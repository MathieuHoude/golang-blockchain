package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	accounts := createWallet().accounts
	goCoin := newBlockchain(2, 10, accounts)
	spew.Dump(goCoin)
}
