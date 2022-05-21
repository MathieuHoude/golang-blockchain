package main

import "crypto/ecdsa"

type Blockchain struct {
	chain                    []*Block
	pendingTransactions      []*Transaction
	difficulty, miningReward int
}

func newBlockchain(_difficulty, _miningReward int, accounts []*Account) *Blockchain {
	var transactions []*Transaction
	for _, account := range accounts {
		transactions = append(transactions, newTransaction("", account.address, 100, nil))
	}
	genesisBlock := newBlock("", _difficulty, transactions)
	return &Blockchain{chain: []*Block{genesisBlock}, difficulty: _difficulty, miningReward: _miningReward}
}

//minePendingTransactions takes all pending transactions and add them to a new block. It then creates a new transaction for the mining reward.
func (b *Blockchain) minePendingTransactions(miningRewardAddress string) {
	latestBlock := b.chain[len(b.chain)-1]
	newBlock := newBlock(latestBlock.hash, b.difficulty, b.pendingTransactions)
	b.chain = append(b.chain, newBlock)
	b.pendingTransactions = []*Transaction{newTransaction("", miningRewardAddress, b.miningReward, nil)}
}

//getBalance calculates the current balance of a provided address.
func (b *Blockchain) getBalance(_address string) int {
	balance := 0

	for _, block := range b.chain {
		for _, transaction := range block.transactions {
			if transaction.Status != "rejected" {
				if transaction.FromAddress == _address {
					balance -= transaction.Amount
				}

				if transaction.ToAddress == _address {
					balance += transaction.Amount
				}
			}

		}
	}

	return balance
}

//submitTransaction creates a new transaction based on the data submitted.
//If the _fromAddress does not have a high enough balance, the transaction is automatically flagged as invalid.
func (b *Blockchain) submitTransaction(_fromAddress, _toAddress string, _amount int, _signingKey *ecdsa.PrivateKey) {
	tx := newTransaction(_fromAddress, _toAddress, _amount, _signingKey)
	balance := b.getBalance(_fromAddress)
	if balance < _amount {
		tx.Status = "invalid"
	}
	b.pendingTransactions = append(b.pendingTransactions, tx)
}

//isValid verifies the integrity of each block in the chain.
func (b *Blockchain) isValid() bool {
	for _, block := range b.chain {
		if !block.isValid() {
			return false
		}
	}

	return true
}
