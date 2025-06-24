package remote

import "context"

type RemoteFn func(ctx context.Context, m map[string]interface{}) (map[string]interface{}, error)




