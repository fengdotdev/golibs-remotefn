package tcpbridge

import (
	"context"
	"fmt"
	"net"
)

func (b *TCPBridge) StartSendingDirect(ctx context.Context, dataToSend []byte) error {

	if b.sending {
		return fmt.Errorf("TCPBridge is already sending")
	}
	b.sending = true
	defer func() {
		b.sending = false
	}()

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", b.remoteAddr, b.outPort)) //TODO FIX FOR IPv6
	if err != nil {
		return fmt.Errorf("error connecting to output port: %v", err)
	}
	defer conn.Close()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		fmt.Printf("Sending data: %s\n", string(dataToSend)) // change
		_, err := conn.Write(dataToSend)
		if err != nil {
			return fmt.Errorf("error writing to connection: %v", err)
		}
	}

	return nil
}
