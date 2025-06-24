package remote

import "context"

type RemoteFn func(ctx context.Context, inMap map[string]interface{}) (map[string]interface{}, error)

type RemoteFetch func(ctx context.Context, inMap map[string]interface{}) ([]byte, error)

type RemoteGOB func(ctx context.Context, gob []byte) ([]byte, error)

type RemoteFetchGOB func(ctx context.Context, gob []byte) ([]byte, error)

type RemoteJSON func(ctx context.Context, json []byte) ([]byte, error)

type RemoteFetchJSON func(ctx context.Context, json []byte) ([]byte, error)
