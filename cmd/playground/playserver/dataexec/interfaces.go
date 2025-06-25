package dataexec

type DataInOutFn func(key string, data []byte) ([]byte, error)

type DataExec interface {
	AddInMiddleware(fn DataInOutFn)
	AddOutMiddleware(fn DataInOutFn)
	DataIn(key string, data []byte) ([]byte, error)
	DataOut(key string, data []byte) ([]byte, error)
	DataInOut(key string, dataIn []byte) ([]byte, error)
}
