package main

import (
	"context"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"
)

func Foo(a, b int) int {
	return a + b
}

func Foohandler(ctx context.Context, m map[string]interface{}) (map[string]interface{}, error) {

	a := remote.Or(m, "a", 0)
	b := remote.Or(m, "b", 0)

	result := Foo(a, b)

	resultMap := remote.ResultSingle(result)

	return resultMap, nil
}

func main() {

	reg := remote.NewGoRegister()

	err := reg.RegisterRemoteFn("Foo", Foohandler)
	if err != nil {
		panic(err)
	}



	


	fn, err := reg.GetRemoteFn("Foo")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	m := map[string]interface{}{
		"a": 5,
		"b": 10,
	}

	result, err := fn(ctx, m)
	if err != nil {
		panic(err)
	}
	if res, ok := result["result"]; ok {
		println("Result:", res.(int))
	}

}
