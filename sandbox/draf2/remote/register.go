package remote

type RegisterRemote interface {
	RegisterRemote_basic
}

type RegisterRemote_basic interface {
	RegisterRemoteFn(name string, fn RemoteFn) error
	GetRemoteFn(name string) (RemoteFn, error)
}

type RegisterRemote_Fetch interface {
	RegisterRemoteFetch(name string, fn RemoteFetch) error
	GetRemoteFetch(name string) (RemoteFetch, error)
}

type RegisterRemote_GOB interface {
	RegisterRemoteGOB(name string, fn RemoteGOB) error
	RegisterRemoteFetchGOB(name string, fn RemoteFetchGOB) error
	GetRemoteGOB(name string) (RemoteGOB, error)
	GetRemoteFetchGOB(name string) (RemoteFetchGOB, error)
}

type RegisterRemote_JSON interface {
	RegisterRemoteJSON(name string, fn RemoteJSON) error
	RegisterRemoteFetchJSON(name string, fn RemoteFetchJSON) error
	GetRemoteJSON(name string) (RemoteJSON, error)
	GetRemoteFetchJSON(name string) (RemoteFetchJSON, error)
}
