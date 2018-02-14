package main

import (
	"fmt"
	"github.com/stellar/go/keypair"
)

func CreateKeypair() (*keypair.Full, error) {
	return keypair.Random()
}

func main() {
	keypair, err := CreateKeypair()
	if err != nil {
		panic(err)
	}

	fmt.Println("Secret seed: ", keypair.Seed())
	fmt.Println("Public Key: ", keypair.Address())
}
