package main

import (
	"fmt"
)

func main() {
	accounts := createWallet().accounts
	goCoin := newBlockchain(4, 50, accounts)
	goCoin.pendingTransactions = append(goCoin.pendingTransactions, newTransaction(accounts[0].address, accounts[1].address, 10, accounts[0].privatekey))
	goCoin.pendingTransactions = append(goCoin.pendingTransactions, newTransaction(accounts[2].address, accounts[3].address, 20, accounts[2].privatekey))
	goCoin.minePendingTransactions(accounts[0].address)
	goCoin.minePendingTransactions(accounts[0].address)
	balance := goCoin.getBalance(accounts[0].address)
	fmt.Printf("Balance of %s: %d \n", accounts[0].address, balance)

}
