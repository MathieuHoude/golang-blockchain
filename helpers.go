package main

func (b *Blockchain) LastTx() *Transaction {
	return b.pendingTransactions[len(b.pendingTransactions)-1]
}

func (b *Blockchain) LastBlock() *Block {
	return b.chain[len(b.chain)-1]
}
