package tcpbridge

import (
	"context"
	"fmt"
	"net"
)

func (b *TCPBridge) StartListening(ctx context.Context) error {

	if b.listening {
		return fmt.Errorf("TCPBridge is already listening")
	}
	b.listening = true
	defer func() {
		b.listening = false
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", b.inPort))
	if err != nil {
		return fmt.Errorf("error starting listener on port %d: %v", b.inPort, err)
	}
	defer listener.Close()

	buffer := make([]byte, 1024)

	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("Error accepting connection: %v\n", err)
		return fmt.Errorf("error accepting connection: %v", err)
	}
	defer conn.Close()

	for {
		select {
		case <-ctx.Done():

			return ctx.Err()
		default:
			fmt.Printf("Listening on port %d...\n", b.inPort)
			length, err := conn.Read(buffer)
			if err != nil {
				return fmt.Errorf("error reading from connection: %v", err)
			}
			received := buffer[:length]
			fmt.Printf("Received data: %s\n", string(received)) // change
		}
	}

}

func (b *TCPBridge) StartListeningNoBlocking(ctx context.Context) (chan []byte, chan error) {

	errChan := make(chan error)

	go func() {
		err := b.StartListening(ctx)
		if err != nil {
			errChan <- err
		}
	}()

	return b.GetInChan(), errChan
}
