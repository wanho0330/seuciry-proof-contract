package chain

import (
	"errors"
	"os"

	"github.com/Netflix/go-env"

	"security-proof-contract/pkg/constants"
)

// Config struct is composed of a PRCUrl, a PrivateKey and a ContractAddress.
type Config struct {
	RPCUrl          string `env:"BLOCK_CHAIN_RPC_URL,default=http://127.0.0.1:50097"`
	PrivateKey      string `env:"BLOCK_CHAIN_PRIVATE_KEY,default=bcdf20249abf0ed6d944c0288fad489e33f66b3960d9e6229c1cd214ed3bbe31"`
	ContractAddress string `env:"BLOCK_CHAIN_CONTRACT_ADDRESS,default=0xb4B46bdAA835F8E4b4d8e208B6559cD267851051"`
}

// InitConfig function is returning a Config and an error.
func InitConfig() (*Config, error) {
	config := &Config{}
	_, err := env.UnmarshalFromEnviron(config)
	if err != nil {
		return nil, errors.Join(constants.ErrConfig, err)
	}

	return &Config{
		RPCUrl:          config.RPCUrl,
		PrivateKey:      config.PrivateKey,
		ContractAddress: config.ContractAddress,
	}, nil
}

// ContractAddress function is returning an error accepting address.
func ContractAddress(address string) error {
	err := os.Setenv("BLOCK_CHAIN_CONTRACT_ADDRESS", address)
	if err != nil {
		return errors.Join(constants.ErrContractAddress, err)
	}

	return nil
}
