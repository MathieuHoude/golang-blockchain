package main

type Transaction struct {
	fromAddress string
	toAddress   string
	amount      int
}

func newTransaction(_fromAddress, _toAddress string, _amount int) *Transaction {
	return &Transaction{fromAddress: _fromAddress, toAddress: _toAddress, amount: _amount}
}
