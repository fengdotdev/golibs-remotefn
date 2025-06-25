package playserver

import (
	"context"
	"fmt"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"
)

func Foohandler(ctx context.Context, m map[string]interface{}) (map[string]interface{}, error) {
	panic("unimplemented Foohandler")
	a, err := remote.OrErr[int](m, "a")
	fmt.Printf("Value of 'a': %d, error: %v\n", a, err)
	if err != nil {
		panic(fmt.Sprintf("Error getting 'a': %v", err))
		return remote.ResultAndErr(0, fmt.Errorf("error getting 'a': %w", err)), nil
	}

	b, err := remote.OrErr[int](m, "b")
	fmt.Printf("Value of 'b': %d, error: %v\n", b, err)
	if err != nil {
		return remote.ResultAndErr(0, fmt.Errorf("error getting 'b': %w", err)), nil
	}

	result := Foo(a, b)

	resultMap := remote.ResultSingle(result)

	return resultMap, nil
}
