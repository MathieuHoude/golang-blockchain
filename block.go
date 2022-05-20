package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	timestamp          time.Time
	transactions       []*Transaction
	previousHash, hash string
	nonce              int
}

func newBlock(_previousHash string, _difficulty int, _pendingTransactions []*Transaction) *Block {
	block := Block{timestamp: time.Now(), previousHash: _previousHash, nonce: 0, transactions: _pendingTransactions}
	block.calculateHash()
	block.mineBlock(_difficulty)
	return &block
}

func (b *Block) calculateHash() string {
	var transactions []byte
	if len(b.transactions) != 0 {
		transactions, _ = json.Marshal(b.transactions)
	} else {
		transactions = []byte("")
	}

	data := []byte(b.timestamp.String() + b.previousHash + string(transactions) + fmt.Sprint(b.nonce))
	hash := sha256.Sum256(data)
	b.hash = hex.EncodeToString(hash[:])
	return b.hash
}

func (b *Block) mineBlock(_difficulty int) {
	for b.hash[0:_difficulty] != strings.Repeat("0", _difficulty) {
		b.calculateHash()
		b.nonce += 1
	}

	fmt.Printf("Block mined: %s \n", b.hash)
}

func (b *Block) isValid(newBlock, oldBlock Block) bool {
	if b.calculateHash() != b.hash {
		return false
	}

	return true
}
