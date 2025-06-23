package main

import (
	"context"
	"time"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draft1/tcpbridge"
)

func main() {

	myTCPBridge, err := tcpbridge.NewTCPBridgeWithReversedA()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dataToSend := []byte("Hello from Play2")
	interval := time.Duration(2) * time.Second

	for {
		select {
		case <-time.After(interval):
			err := myTCPBridge.StartSendingDirect(ctx, dataToSend)
			if err != nil {
				println("Error sending data:", err.Error())
				cancel() // Cancel context on error to stop the loop
			} else {
				println("data was send:", string(dataToSend))
			}
		case <-ctx.Done():
			println("Context cancelled, stopping...")
			return
		}
	}

}
