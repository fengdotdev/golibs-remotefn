package tcpbridge

func (b *TCPBridge) Close() {
	close(b.inChan)
	close(b.outChan)
}

func (b *TCPBridge) Reset() {
	b.inChan = make(chan []byte)
	b.outChan = make(chan []byte)
	b.onlyOnceOut = false
	b.onlyOnceIn = false
}
