// Package constants is a package for handling constants.
package constants

import "errors"

// Defines errors related to the chain service.
var (
	ErrNewClient          = errors.New("new client error")
	ErrPrePareAuth        = errors.New("prepares auth error")
	ErrConfig             = errors.New("config error")
	ErrContractAddress    = errors.New("contract address error")
	ErrDeploy             = errors.New("deploy error")
	ErrConfirmProof       = errors.New("confirm proof error")
	ErrConfirmUpdateProof = errors.New("confirm update proof error")
	ErrReadImageHashes    = errors.New("read image hashes error")
	ErrReadLastImageHash  = errors.New("read last image hash error")
)
