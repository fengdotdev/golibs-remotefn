package tcpbridge

import (
	"fmt"
	"sync"
)

const (
	defaultInPort  = 8080
	defaultOutPort = 8081
	a_InPort       = 8082 //
	a_OutPort      = 8084
	b_InPort       = 8085
	b_OutPort      = 8086
	queueSize      = 10
)

var (
	instance *TCPBridge
	once     sync.Once
)

// NewTCPBridge creates a new instance of TCPBridge with the specified input and output ports.
// It uses sync.Once to ensure that only one instance is created.

func NewTCPBridgeA() (*TCPBridge, error) {
	return NewTCPBridge(a_InPort, a_OutPort)
}

func NewTCPBridgeWithReversedA() (*TCPBridge, error) {
	return NewTCPBridge(a_OutPort, a_InPort)
}

func NewTCPBridgeB() (*TCPBridge, error) {
	return NewTCPBridge(b_InPort, b_OutPort)
}

func NewTCPBridgeWithDefaults() (*TCPBridge, error) {
	return NewTCPBridge(defaultInPort, defaultOutPort)
}

func NewTCPBridge(inPort, outPort int) (*TCPBridge, error) {
	var err error
	once.Do(func() {
		instance = &TCPBridge{
			inPort:       inPort,
			outPort:      outPort,
			inChan:       make(chan []byte),
			outChan:      make(chan []byte),
			sendingQueue: make([]chan []byte, queueSize),
			remoteAddr:   "localhost", // Default remote address
		}
	})
	if instance == nil {
		return nil, fmt.Errorf("TCPBridge instance already created")
	}
	return instance, err
}
