package remote

import (
	"context"
	"fmt"
)

func MockWire(ctx context.Context, datachan chan []byte, key string, fn func(key string, data []byte) []byte) chan []byte {

	output := make(chan []byte)
	go func(key string) {
		defer close(output)

		for {
			select {
			case data := <-datachan:
				if data == nil {
					return
				}
				fmt.Printf("Received data at mock for key %s: %s\n", key, string(data))
				result := fn(key, data)
				fmt.Println("Result from mock function:", string(result))
				if result == nil {
					return
				}

				output <- result
			case <-ctx.Done():
				return
			}
		}
	}(key)
	return output
}
