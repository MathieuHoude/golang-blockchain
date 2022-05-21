package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Transaction struct {
	FromAddress string `json:"fromAddress"`
	ToAddress   string `json:"toAddress"`
	Status      string `json:"status"`
	Amount      int    `json:"amount"`
	Signature   []byte `json:"signature"`
}

func newTransaction(_fromAddress, _toAddress string, _amount int, _signingKey *ecdsa.PrivateKey) *Transaction {
	tx := &Transaction{FromAddress: _fromAddress, ToAddress: _toAddress, Amount: _amount}
	if _signingKey != nil {
		tx.signTransaction(_signingKey)
	}
	return tx
}

func (t *Transaction) calculateHash() common.Hash {
	data := []byte(t.FromAddress + t.ToAddress + fmt.Sprint(t.Amount))
	hash := crypto.Keccak256Hash(data)
	return hash
}

func (t *Transaction) signTransaction(signingKey *ecdsa.PrivateKey) {
	hash := t.calculateHash()
	signature, err := crypto.Sign(hash.Bytes(), signingKey)
	if err != nil {
		log.Fatal(err)
	}
	t.Signature = signature
}

func (t *Transaction) isValid() bool {

	if t.FromAddress == "" {
		return true
	}

	if t.Signature == nil || len(t.Signature) == 0 || t.Status == "invalid" {
		return false
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(t.calculateHash().Bytes(), t.Signature)
	if err != nil {
		log.Fatal(err)
	}
	address := crypto.PubkeyToAddress(*sigPublicKeyECDSA).Hex()

	if t.FromAddress != address {
		return false
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	signatureNoRecoverID := t.Signature[:len(t.Signature)-1]
	verified := crypto.VerifySignature(sigPublicKeyBytes, t.calculateHash().Bytes(), signatureNoRecoverID)

	return verified
}
