// Package main is the server for running dApp.
package main

import (
	"net/http"
	"time"

	"buf.build/gen/go/wanho/security-proof-api/connectrpc/go/chain/v1/chainv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"security-proof-contract/internal/chain"
)

func main() {
	baseAddr := "127.0.0.4:8090"

	config, err := chain.InitConfig()
	if err != nil {
		panic(err)
	}

	client, auth := chain.Init(config)
	proof := chain.NewHandler(config.ContractAddress, client, auth)

	mux := http.NewServeMux()
	path, handler := chainv1connect.NewProofServiceHandler(proof)
	mux.Handle(path, handler)

	server := &http.Server{
		Addr:              baseAddr,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}
	_ = server.ListenAndServe()
}
