package remote

type RegisterRemote interface {
	RegisterRemoteFn(name string, fn RemoteFn) error
	GetRemoteFn(name string) (RemoteFn, error)
}
