package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	accounts []*Account
}

type Account struct {
	privatekey      *ecdsa.PrivateKey
	privateKeyBytes []byte
	publicKey       *ecdsa.PublicKey
	publicKeyBytes  []byte
	address         string
}

func createWallet() *Wallet {
	var accounts []*Account
	//Create 10 accounts
	for i := 0; i < 10; i++ {
		var account *Account = createAccount()
		accounts = append(accounts, account)
	}

	return &Wallet{accounts: accounts}
}

func createAccount() *Account {
	privatekey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // this generates a public & private key pair
	privateKeyBytes := crypto.FromECDSA(privatekey)
	publicKey := privatekey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return &Account{privatekey: privatekey, privateKeyBytes: privateKeyBytes, publicKey: publicKeyECDSA, publicKeyBytes: publicKeyBytes, address: address}

}
