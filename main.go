package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	accounts := createWallet().accounts
	goCoin := newBlockchain(2, 10, accounts)
	goCoin.submitTransaction(accounts[0].address, accounts[1].address, 10, accounts[0].privatekey)
	goCoin.submitTransaction(accounts[2].address, accounts[3].address, 20, accounts[2].privatekey)
	goCoin.minePendingTransactions(accounts[0].address)
	goCoin.minePendingTransactions(accounts[0].address)
	valid := goCoin.isValid()
	fmt.Printf("%v \n", valid)
	goCoin.chain[0].transactions[0].Amount = 50
	valid = goCoin.isValid()
	fmt.Printf("%v \n", valid)
	spew.Dump(goCoin)
}
