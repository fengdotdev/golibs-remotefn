package tcpbridge

import (
	"context"
)

type TCPBridge struct {
	inPort       int
	outPort      int
	listening    bool
	sending      bool
	sendingQueue []chan []byte // channels to send data
	inChan       chan []byte   // block until data is received
	outChan      chan []byte   // block until data is sent
	onlyOnceOut  bool          // ensure that the output channel is only used once
	onlyOnceIn   bool          // ensure that the input channel is only used once
	ctx          context.CancelFunc
	remoteAddr   string // remote address for sending data
}
