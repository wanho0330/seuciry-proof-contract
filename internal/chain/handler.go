package chain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	chainv1 "buf.build/gen/go/wanho/security-proof-api/protocolbuffers/go/chain/v1"
	"connectrpc.com/connect"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"security-proof-contract/pkg/constants"
	"security-proof-contract/pkg/contract/proof"
)

// Handler struct is composed of a contractAddress, a Client and a TransactOpts.
type Handler struct {
	contractAddress string
	client          *ethclient.Client
	auth            *bind.TransactOpts
}

// NewHandler function is returning a Handler, accepting a contractAddress, a Client and TransactOpts.
func NewHandler(contractAddress string, client *ethclient.Client, auth *bind.TransactOpts) *Handler {
	return &Handler{contractAddress: contractAddress, client: client, auth: auth}
}

// ConfirmProof method is returning a ConfirmProofResponse and an error, accepting a context and ConfirmProofRequest
func (h *Handler) ConfirmProof(ctx context.Context, req *connect.Request[chainv1.ConfirmProofRequest]) (*connect.Response[chainv1.ConfirmProofResponse], error) {
	instance, err := newInstance(h.contractAddress, h.client)
	if err != nil {
		return nil, errors.Join(constants.ErrConfirmProof, err)
	}

	tx, err := instance.ConfirmProof(h.auth, big.NewInt(int64(req.Msg.Idx)), req.Msg.FirstImageHash, req.Msg.SecondImageHash)
	if err != nil {
		return nil, errors.Join(constants.ErrConfirmProof, err)
	}

	var receipt *types.Receipt
	for i := 0; i < 10; i++ { // 최대 10번 반복
		receipt, err = h.client.TransactionReceipt(ctx, tx.Hash())
		if err == nil { // 트랜잭션이 성공적으로 확인되면 종료
			break
		}
		time.Sleep(2 * time.Second) // 2초 대기
	}
	if err != nil {
		return nil, errors.Join(constants.ErrConfirmProof, err)
	}

	var tokenId int64

	for i, log := range receipt.Logs {
		event, err := instance.ParseProofConfirmed(*log)
		fmt.Println(i, err) // err가 발생하면 출력해주는 로직 (정지 안됨)
		if event != nil {
			tokenId = event.TokenId.Int64() // Assign token ID from event
			break
		}
	}

	res := connect.NewResponse(&chainv1.ConfirmProofResponse{
		TokenId: int32(tokenId),
	})

	return res, nil
}

// ConfirmUpdateProof method is returning a ConfirmUpdateProofResponse and an error, accepting a context and ConfirmUpdateProofRequest
func (h *Handler) ConfirmUpdateProof(_ context.Context, req *connect.Request[chainv1.ConfirmUpdateProofRequest]) (*connect.Response[chainv1.ConfirmUpdateProofResponse], error) {
	instance, err := newInstance(h.contractAddress, h.client)
	if err != nil {
		return nil, errors.Join(constants.ErrConfirmUpdateProof, err)
	}

	_, err = instance.ConfirmUpdateProof(h.auth, big.NewInt(int64(req.Msg.TokenId)), req.Msg.FirstImageHash, req.Msg.SecondImageHash)
	if err != nil {
		return nil, errors.Join(constants.ErrConfirmUpdateProof, err)
	}

	res := &connect.Response[chainv1.ConfirmUpdateProofResponse]{}
	return res, nil
}

// ReadImageHashes method is returning a ReadImageHashesResponse and an error, accepting a context and ReadImageHashesRequest
func (h *Handler) ReadImageHashes(_ context.Context, req *connect.Request[chainv1.ReadImageHashesRequest]) (*connect.Response[chainv1.ReadImageHashesResponse], error) {

	instance, err := newInstance(h.contractAddress, h.client)
	if err != nil {
		return nil, errors.Join(constants.ErrReadImageHashes, err)
	}

	tokenId := big.NewInt(int64(req.Msg.TokenId))
	firstImageHashes, secondImageHashes, err := instance.ReadImageHashes(&bind.CallOpts{}, tokenId)
	if err != nil {
		return nil, errors.Join(constants.ErrReadImageHashes, err)
	}

	res := connect.NewResponse(&chainv1.ReadImageHashesResponse{
		FirstImageHashes:  firstImageHashes,
		SecondImageHashes: secondImageHashes,
	})

	return res, nil
}

// ReadLastImageHash method is returning a ReadLastImageHashResponse and an error, accepting a context and ReadLastImageHashRequest
func (h *Handler) ReadLastImageHash(_ context.Context, req *connect.Request[chainv1.ReadLastImageHashRequest]) (*connect.Response[chainv1.ReadLastImageHashResponse], error) {

	instance, err := newInstance(h.contractAddress, h.client)
	if err != nil {
		return nil, errors.Join(constants.ErrReadLastImageHash, err)
	}

	tokenId := big.NewInt(int64(req.Msg.TokenId))
	firstImageHash, secondImageHash, err := instance.ReadLatestImageHash(&bind.CallOpts{}, tokenId)
	if err != nil {
		return nil, errors.Join(constants.ErrReadLastImageHash, err)
	}

	res := connect.NewResponse(&chainv1.ReadLastImageHashResponse{
		FirstImageHash:  firstImageHash,
		SecondImageHash: secondImageHash,
	})

	return res, nil
}

func newInstance(contractAddress string, client *ethclient.Client) (*proof.Proof, error) {
	proofAddress := common.HexToAddress(contractAddress)
	instance, err := proof.NewProof(proofAddress, client)
	if err != nil {
		return nil, errors.Join(constants.ErrConfirmProof, err)
	}

	return instance, nil
}
