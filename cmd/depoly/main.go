// Package main is the server for running chain deploy.
package main

import (
	"fmt"
	"log"

	"security-proof-contract/internal/chain"
)

func main() {
	config, err := chain.InitConfig()
	if err != nil {
		log.Fatalf("Failed to config block chain: %s", err)
	}

	client, auth := chain.Init(config)

	address, tx, _, err := chain.Deploy(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy BlogPostNFT: %s", err)
	}

	err = chain.ContractAddress(address.Hex())
	if err != nil {
		log.Fatalf("Failed to set contract address: %s", err)
	}

	fmt.Printf("Contract deployed to: %s\nTransaction: %s\n", address.Hex(), tx.Hash().Hex())
}
