package main

import (
	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	"os"
)

func main() {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	var app types.Application = NewKVStoreApplication()

	// Start the listener
	srv, err := server.NewServer("tcp://0.0.0.0:26658", "grpc", app)
	if err != nil {
		panic(err)
	}
	srv.SetLogger(logger.With("module", "abci-server"))
	if err := srv.Start(); err != nil {
		panic(err)
	}

	// Wait forever
	cmn.TrapSignal(func() {
		// Cleanup
		srv.Stop()
	})
}
