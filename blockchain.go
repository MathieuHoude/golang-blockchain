package main

type Blockchain struct {
	chain               []*Block
	pendingTransactions []*Transaction
	difficulty          int
	miningReward        int
}

func newBlockchain(_difficulty, _miningReward int) *Blockchain {
	genesisBlock := newBlock("", _difficulty, []*Transaction{newTransaction("", "myaddress", 100, nil), newTransaction("", "myaddress2", 100, nil)})
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
			if transaction.fromAddress == _address {
				balance -= transaction.amount
			}

			if transaction.toAddress == _address {
				balance += transaction.amount
			}
		}
	}

	return balance
}
