package main

import (
	"context"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"
)

func Foo(a, b int) int {
	return a + b
}

func OrErr[T any](m map[string]interface{}, key string) (T, error) {
	var zero T

	if value, ok := m[key]; ok {
		return value.(T), nil
	}
	return zero, nil
}

func Or[T any](m map[string]interface{}, key string, defaultValue T) T {
	if value, ok := m[key]; ok {
		return value.(T)
	}
	return defaultValue
}

func ResultSingle[T any](value T) map[string]interface{} {
	return map[string]interface{}{
		"result": value,
	}
}

func ResultAndErr[T any](value T, err error) map[string]interface{} {
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}
	return map[string]interface{}{
		"result": value,
	}
}

func Foohandler(ctx context.Context, m map[string]interface{}) (map[string]interface{}, error) {

	a := Or(m, "a", 0)
	b := Or(m, "b", 0)

	result := Foo(a, b)

	resultMap := ResultSingle(result)

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
