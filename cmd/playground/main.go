package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"
)

func op(key string, data []byte) []byte {

	errmap := remote.ResultAndErr(1, errors.New("some error"))
	m, err := remote.DataToMapOnJson(data)
	if err != nil {
		outdata, err := remote.MapToDataOnJson(errmap)
		if err != nil {
			panic(fmt.Sprintf("Error converting errlr to map: %v", err))
		}
		return outdata
	}

	result, err := Foohandler(context.Background(), m)
	if err != nil {
		outdata, err := remote.MapToDataOnJson(remote.ResultAndErr(0, err))
		if err != nil {
			panic(fmt.Sprintf("Error converting error to data: %v", err))
		}
		return outdata
	}
	outdata, err := remote.MapToDataOnJson(result)
	if err != nil {
		panic(fmt.Sprintf("Error converting result to data: %v", err))
	}
	return outdata
}

func FooRemote(a, b int) (int, error) {
	ctx := context.Background()
	m := map[string]interface{}{
		"a": a,
		"b": b,
	}
	data, err := remote.MapToDataOnJson(m)
	if err != nil {
		return 1, err
	}
	dataChan := make(chan []byte)

	returndata := remote.MockWire(ctx, dataChan, "Foo", op)

	dataChan <- data
	select {
	case result := <-returndata:
		resultMap, err := remote.DataToMapOnJson(result)
		if err != nil {
			return 0, fmt.Errorf("error converting data to map: %v", err)
		}
		if res, ok := resultMap["result"]; ok {

			// Check if the result is of type int
			if res == nil {
				return 0, fmt.Errorf("result is nil")
			}
			if _, ok := res.(int); !ok {
				return 0, fmt.Errorf("result is not of type int, got %T", res)
			}

			return res.(int), nil
		}

		return 1, nil
	}
}

func Foo(a, b int) int {
	return a + b
}

func Foohandler(ctx context.Context, m map[string]interface{}) (map[string]interface{}, error) {

	a, err := remote.OrErr[int](m, "a")
	if err != nil {
		return remote.ResultAndErr(0, fmt.Errorf("error getting 'a': %w", err)), nil
	}

	b, err := remote.OrErr[int](m, "b")
	if err != nil {
		return remote.ResultAndErr(0, fmt.Errorf("error getting 'b': %w", err)), nil
	}

	result := Foo(a, b)

	resultMap := remote.ResultSingle(result)

	return resultMap, nil
}

func main() {

	result, err := FooRemote(1, 2)
	if err != nil {
		panic(fmt.Sprintf("Error in FooRemote: %v", err))
	}
	println("Result:", result)

}
