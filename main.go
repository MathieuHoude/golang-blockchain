package main

import "fmt"

func main() {
	goCoin := newBlockchain(4, 50)
	goCoin.pendingTransactions = append(goCoin.pendingTransactions, newTransaction("11", "22", 10))
	goCoin.pendingTransactions = append(goCoin.pendingTransactions, newTransaction("22", "11", 1))
	goCoin.minePendingTransactions("myaddress")
	goCoin.minePendingTransactions("myaddress")
	balance := goCoin.getBalance("myaddress")
	fmt.Printf("%d", balance)
}
