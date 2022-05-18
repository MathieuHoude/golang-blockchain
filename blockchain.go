package main

import (
	"time"
)

type Blockchain struct {
	chain               []*Block
	pendingTransactions []*Transaction
	difficulty          int
	miningReward        int
}

func newBlockchain(_difficulty, _miningReward int) *Blockchain {
	genesisBlock := &Block{timestamp: time.Now(), previousHash: "", transactions: []*Transaction{}}
	return &Blockchain{chain: []*Block{genesisBlock}, difficulty: _difficulty, miningReward: _miningReward}
}

func (b *Blockchain) minePendingTransactions(miningRewardAddress string) {
	latestBlock := b.chain[len(b.chain)-1]
	newBlock := newBlock(latestBlock.hash, b.difficulty, b.pendingTransactions)
	b.chain = append(b.chain, newBlock)
	b.pendingTransactions = []*Transaction{newTransaction("", "myaddress", b.miningReward)}
}

func (b *Blockchain) getBalance(_address string) int {
	balance := 0

	for _, transaction := range b.pendingTransactions {
		if transaction.fromAddress == _address {
			balance -= transaction.amount
		}

		if transaction.toAddress == _address {
			balance += transaction.amount
		}
	}

	return balance
}
