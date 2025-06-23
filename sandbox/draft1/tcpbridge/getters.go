package tcpbridge

func (b *TCPBridge) GetInPort() int {
	return b.inPort
}

func (b *TCPBridge) GetOutPort() int {
	return b.outPort
}

func (b *TCPBridge) GetInChan() chan []byte {

	if b.onlyOnceIn {
		panic("Input channel can only be used once")
	}
	b.onlyOnceIn = true

	return b.inChan
}

func (b *TCPBridge) GetOutChan() chan []byte {

	if b.onlyOnceOut {
		panic("Output channel can only be used once")
	}
	b.onlyOnceOut = true

	return b.outChan
}
