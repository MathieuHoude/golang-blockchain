package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

type Transaction struct {
	fromAddress string
	toAddress   string
	amount      int
	signature   []byte
}

func newTransaction(_fromAddress, _toAddress string, _amount int, _signingKey *ecdsa.PrivateKey) *Transaction {
	tx := &Transaction{fromAddress: _fromAddress, toAddress: _toAddress, amount: _amount}
	if _signingKey != nil {
		tx.signTransaction(_signingKey)
	}
	return tx
}

func (t *Transaction) calculateHash() []byte {
	data := []byte(t.fromAddress + t.toAddress + fmt.Sprint(t.amount))
	hash := sha256.Sum256(data)
	return hash[:]
}

func (t *Transaction) signTransaction(signingKey *ecdsa.PrivateKey) {
	hash := t.calculateHash()
	r, s, _ := ecdsa.Sign(rand.Reader, signingKey, hash)
	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)
	t.signature = signature
}

func (t *Transaction) isValid() bool {
	if t.fromAddress == "" {
		return true
	}

	if t.signature == nil || len(t.signature) == 0 {
		return false
	}

	// Verify
	verifystatus := ecdsa.Verify(&signingKey.PublicKey, t.calculateHash(), r, s)
	fmt.Println(verifystatus) // should be true
}
