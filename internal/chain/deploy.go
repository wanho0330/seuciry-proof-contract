package chain

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"security-proof-contract/pkg/constants"
	"security-proof-contract/pkg/contract/proof"
)

// Deploy function is returning an Address, a Transaction, a Proof and an error, accepting a TransactOpts and a Client.
func Deploy(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *types.Transaction, *proof.Proof, error) {
	address, tx, instance, err := proof.DeployProof(auth, client)
	if err != nil {
		return [20]byte{}, nil, nil, errors.Join(constants.ErrDeploy, err)
	}

	return address, tx, instance, nil

}
