package handlers

import (
	"context"
	"fmt"

	"github.com/fengdotdev/golibs-remotefn/cmd/playground/playserver/funcs"
	"github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"
)

func AddHandler(ctx context.Context, m map[string]interface{}) (map[string]interface{}, error) {

	a, err := remote.OrErr[int](m, "a")
	if err != nil {
		return remote.ResultAndErr(0, fmt.Errorf("error getting 'a': %w", err)), nil
	}

	b, err := remote.OrErr[int](m, "b")
	if err != nil {
		return remote.ResultAndErr(0, fmt.Errorf("error getting 'b': %w", err)), nil
	}

	result := funcs.Add(a, b)

	resultMap := remote.ResultSingle(result)

	return resultMap, nil
}
