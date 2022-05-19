package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func main() {
	// priv, x, y, err := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	// x := priv.PublicKey
	privatekey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // this generates a public & private key pair
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("Private Key :")
	fmt.Printf("%x \n", privatekey)
	fmt.Println("Public Key :")
	fmt.Printf("%x \n", privatekey.PublicKey)
	goCoin := newBlockchain(4, 50)
	goCoin.pendingTransactions = append(goCoin.pendingTransactions, newTransaction("11", "22", 10, privatekey))
	goCoin.pendingTransactions = append(goCoin.pendingTransactions, newTransaction("22", "11", 1, privatekey))
	goCoin.minePendingTransactions("myaddress")
	goCoin.minePendingTransactions("myaddress")
	balance := goCoin.getBalance("myaddress")
	fmt.Printf("%d", balance)

}
