package main

import (
	"context"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draft1/tcpbridge"
)

func main() {

	myTCPBridge, err := tcpbridge.NewTCPBridgeA()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err = myTCPBridge.StartListening(ctx)
	if err != nil {
		panic(err)
	}
}
