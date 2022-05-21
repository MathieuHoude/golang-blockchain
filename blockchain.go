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

func (b *Blockchain) minePendingTransactions(miningRewardAddress string) {
	latestBlock := b.chain[len(b.chain)-1]
	newBlock := newBlock(latestBlock.hash, b.difficulty, b.pendingTransactions)
	b.chain = append(b.chain, newBlock)
	b.pendingTransactions = []*Transaction{newTransaction("", miningRewardAddress, b.miningReward, nil)}
}

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

func (b *Blockchain) submitTransaction(_fromAddress, _toAddress string, _amount int, _signingKey *ecdsa.PrivateKey) {
	tx := newTransaction(_fromAddress, _toAddress, _amount, _signingKey)
	balance := b.getBalance(_fromAddress)
	if balance < _amount {
		tx.Status = "invalid"
	}
	b.pendingTransactions = append(b.pendingTransactions, tx)
}

func (b *Blockchain) isValid() bool {
	for _, block := range b.chain {
		if !block.isValid() {
			return false
		}
	}

	return true
}
