package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//Transaction contains the details of a transfer between two accounts, along with the signature of it's creator
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

//calculateHash hashes the transaction's content
func (t *Transaction) calculateHash() common.Hash {
	data := []byte(t.FromAddress + t.ToAddress + fmt.Sprint(t.Amount))
	hash := crypto.Keccak256Hash(data)
	return hash
}

//signTransaction creates a hash based on the transaction's content and the private key of the transaction creator. It allows us to
func (t *Transaction) signTransaction(signingKey *ecdsa.PrivateKey) {
	hash := t.calculateHash()
	signature, err := crypto.Sign(hash.Bytes(), signingKey)
	if err != nil {
		log.Fatal(err)
	}
	t.Signature = signature
}

//isValid verify the integrity of a transaction.
//If there is no FromAddress, the transaction refers to mining reward and does not need additional checks.
//We then make sure the transaction has a signature and was not already flagged as invalid.
//Finally, we verify that the key which signed the transaction matches with the sender address.
func (t *Transaction) isValid() bool {

	if t.Status == "invalid" {
		return false
	}

	if t.FromAddress == "" {
		return true
	}

	if t.Signature == nil || len(t.Signature) == 0 {
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
