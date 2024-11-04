// Package chain is a package for handling blockchain processes.
package chain

import (
	"context"
	"errors"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"security-proof-contract/pkg/constants"
)

// Init function is returning a Client and TransactOpts, accepting Config
func Init(config *Config) (*ethclient.Client, *bind.TransactOpts) {

	client, err := newClient(config.RPCUrl)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	auth, err := prePareAuth(client, config.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to prepare auth: %v", err)
	}

	return client, auth
}

func newClient(rpcUrl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, errors.Join(constants.ErrNewClient, err)
	}
	return client, nil
}

func prePareAuth(client *ethclient.Client, privateKey string) (*bind.TransactOpts, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.Join(constants.ErrPrePareAuth, err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, errors.Join(constants.ErrPrePareAuth, err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, errors.Join(constants.ErrPrePareAuth, err)
	}

	return auth, nil
}
