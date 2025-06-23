package tcpbridge

import (
	"context"
	"fmt"
	"io"
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

	// Accept connections in a loop
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Printf("Listening on port %d...\n", b.inPort)
			conn, err := listener.Accept()
			if err != nil {
				// Handle temporary errors (e.g., network interruptions) gracefully
				if ctx.Err() != nil {
					return ctx.Err()
				}
				fmt.Printf("Error accepting connection: %v\n", err)
				continue
			}

			// Handle each connection in a separate goroutine
			go b.handleConnection(ctx, conn)
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

func (b *TCPBridge) handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			length, err := conn.Read(buffer)
			if err != nil {
				if err == io.EOF {
					fmt.Printf("Connection closed by client\n")
					return // Client closed the connection, exit this goroutine
				}
				fmt.Printf("Error reading from connection: %v\n", err)
				return // Other errors, exit this goroutine
			}
			received := buffer[:length]
			fmt.Printf("Received data: %s\n", string(received))
			b.inChan <- received // Send data to inChan
		}
	}
}
