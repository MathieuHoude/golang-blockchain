package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

// Create a new Public-Private ECC-256 Keypair.
func CreateKey(log chan string) ([]byte, *big.Int, *big.Int) {
	priv, x, y, err := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log <- "Key Generation Error"
		return nil, nil, nil
	}
	return priv, x, y
}
