package main

import (
	"encoding/json"
	"testing"
)

var accounts = createWallet().accounts
var difficulty int = 2
var miningReward int = 10
var goCoin *Blockchain = newBlockchain(difficulty, miningReward, accounts)

func TestNewBlockchain(t *testing.T) {
	//Check if chain's attributes are properly assigned
	if goCoin.difficulty != difficulty || goCoin.miningReward != miningReward {
		t.Fatalf("Blockchain is not properly initialized. \n Difficulty expected: %v, got %v. \n Mining reward expected: %v, got %v.", difficulty, goCoin.difficulty, miningReward, goCoin.miningReward)
	}

	//Check if chain contains the Genesis block
	if len(goCoin.chain) != 1 {
		t.Fatalf("Incorrect amount of block, expected 1, got %v", len(goCoin.chain))
	}

	//Check if Genesis block is valid
	if !goCoin.chain[0].isValid() {
		t.Fatalf("Genesis block is invalid")
	}

}

func TestNewTransaction(t *testing.T) {
	amountOfTransactionsBefore := len(goCoin.pendingTransactions)
	goCoin.submitTransaction(accounts[0].address, accounts[1].address, 10, accounts[0].privatekey)

	//Check if new transaction was added the to chain's pending transactions
	if len(goCoin.pendingTransactions) != amountOfTransactionsBefore+1 {
		t.Fatalf("Incorrect amount of pending transactions. Expected %v, got %v.", amountOfTransactionsBefore+1, len(goCoin.pendingTransactions))
	}

	//Check if transaction's attributes are properly assigned
	tx := goCoin.pendingTransactions[0]
	if tx.FromAddress != accounts[0].address || tx.ToAddress != accounts[1].address || tx.Amount != 10 {
		t.Fatalf("Transaction's attribute are not properly assigned. \n FromAddress expected: %v, got %v. \n ToAddress expected: %v, got %v. \n Amount expected: %v, got %v.", accounts[0].address, tx.FromAddress, accounts[1].address, tx.ToAddress, 10, tx.Amount)
	}

	//Check if the new transaction is valid
	if !goCoin.pendingTransactions[0].isValid() {
		t.Fatalf("Test transaction is invalid")
	}
}

func TestInvalidTransaction(t *testing.T) {
	//Check if we can submit a transaction with null FromAddress, simulating a mining reward tx
	goCoin.submitTransaction("", accounts[1].address, 10, accounts[1].privatekey)
	if goCoin.LastTx().isValid() {
		t.Fatalf("Fake mining reward transaction should be invalidated")
	}

	//Check if we can submit a transaction in somebody else's name
	goCoin.submitTransaction(accounts[0].address, accounts[1].address, 10, accounts[1].privatekey)
	if goCoin.LastTx().isValid() {
		t.Fatalf("A transaction signed with an account that is not the FromAddress should be rejected")
	}

	//Check if a transaction where the FromAddress has insufficient balance gets rejected
	goCoin.submitTransaction(accounts[0].address, accounts[1].address, 10000, accounts[1].privatekey)
	if goCoin.LastTx().isValid() {
		t.Fatalf("A transaction where the FromAddress has insufficient balance should be rejected")
	}

	//All invalid transactions should be included in a new block and marked as rejected
	goCoin.minePendingTransactions(accounts[0].address)
	txs := goCoin.LastBlock().transactions[1:] //We remove the mining reward tx
	if len(txs) != 3 {
		t.Fatalf("Incorrect amount of rejected transactions in new Block. Expected %v, got %v.", 3, len(txs))
	}
	for _, tx := range txs {
		if tx.Status != "rejected" {
			txJSON, _ := json.Marshal(tx)
			t.Fatalf("Transaction %s should have been marked as rejected", txJSON)
		}
	}
}

func TestNewBlock(t *testing.T) {
	amountOfBlocksBefore := len(goCoin.chain)
	goCoin.submitTransaction(accounts[0].address, accounts[1].address, 10, accounts[0].privatekey)
	goCoin.submitTransaction(accounts[2].address, accounts[3].address, 10, accounts[2].privatekey)
	goCoin.submitTransaction(accounts[4].address, accounts[5].address, 10, accounts[4].privatekey)

	pendingTransactionsAmount := len(goCoin.pendingTransactions)

	goCoin.minePendingTransactions(accounts[9].address)

	//Check if new Block was added to the chain
	if len(goCoin.chain) != amountOfBlocksBefore+1 {
		t.Fatalf("Incorrect amount of blocks. Expected %v, got %v.", amountOfBlocksBefore+1, len(goCoin.chain))
	}

	newBlock := goCoin.chain[len(goCoin.chain)-1]

	//Check if block includes all pending transactions
	if len(newBlock.transactions) != pendingTransactionsAmount {
		t.Fatalf("Incorrect amount of transactions in block. Expected %v, got %v.", pendingTransactionsAmount, len(newBlock.transactions))
	}

	//Check if block's previousHash matches with the previous block's hash
	previousBlock := goCoin.chain[len(goCoin.chain)-2]
	if goCoin.LastBlock().previousHash != previousBlock.hash {
		t.Fatalf("New block's previousHash does not match with the previous block's hash. Expected %v, got %v.", previousBlock.hash, goCoin.LastBlock().previousHash)
	}

	//Check if a mining reward transaction was created
	if len(goCoin.pendingTransactions) != 1 {
		t.Fatalf("No mining reward transaction created after new block")
	}

	miningRewardTx := goCoin.pendingTransactions[0]
	//Check if mining reward transaction has the proper attributes
	if miningRewardTx.FromAddress != "" || miningRewardTx.ToAddress != accounts[9].address || miningRewardTx.Amount != miningReward {
		t.Fatalf("Transaction's attribute are not properly assigned. \n FromAddress expected: %v, got %v. \n ToAddress expected: %v, got %v. \n Amount expected: %v, got %v.", "\"\"", miningRewardTx.FromAddress, accounts[9].address, miningRewardTx.ToAddress, miningReward, miningRewardTx.Amount)
	}
}
