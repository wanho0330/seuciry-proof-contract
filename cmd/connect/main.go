// Package main is the server for running connect test.
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"

	"security-proof-contract/internal/chain"
)

func main() {
	load, err := chain.InitConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client, _ := chain.Init(load)

	privateKeys := []string{
		"bcdf20249abf0ed6d944c0288fad489e33f66b3960d9e6229c1cd214ed3bbe31",
		"53321db7c1e331d93a11a41d16f004d7ff63972ec8ec7c25db329728ceeb1710",
		"ab63b23eb7941c1251757e24b3d2350d2bc05c3c388d06f8fe6feafefb1e8c70",
		"5d2344259f42259f82d2c140aa66102ba89b57b4883ee441a8b312622bd42491",
		"27515f805127bebad2fb9b183508bdacb8c763da16f54e0678b16e8f28ef3fff",
		"7ff1a4c1d57e5e784d327c4c7651e952350bc271f156afb3d00d20f5ef924856",
	}

	for _, pk := range privateKeys {
		privateKey, err := crypto.HexToECDSA(pk)
		if err != nil {
			log.Fatalf("Failed to parse private key: %v", err)
		}

		publicKey := privateKey.Public().(*ecdsa.PublicKey)
		address := crypto.PubkeyToAddress(*publicKey)

		balance, err := client.BalanceAt(context.Background(), address, nil)
		if err != nil {
			log.Fatalf("Failed to get balance: %v", err)
		}

		fmt.Printf("%s has balance %s\n", address.Hex(), balance.String())
	}
}
