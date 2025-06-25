package playclient

import (
	"fmt"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"
)

func FooRemote(a, b int) (int, error) {
	m := map[string]interface{}{
		"a": a,
		"b": b,
	}
	value, err := PipeSingleResult[int]("Foo", m)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func PipeSingleResult[T any](key string, m map[string]interface{}) (T, error) {
	var zero T
	data, err := remote.MapToDataOnJson(m)
	if err != nil {
		return zero, err
	}

	exec := NewDataExec()
	outdata, err := exec.DataInOut(key, data)
	if err != nil {
		return zero, err
	}
	result, err := remote.DataToMapOnJson(outdata)
	if err != nil {
		return zero, err
	}
	fmt.Printf("Result map: %v\n", result)

	value, err := remote.OrErr[T](result, "result")
	if err != nil {
		return zero, err
	}
	fmt.Printf("Value of 'result': %v, error: %v\n", value, err)
	return value, nil

}
