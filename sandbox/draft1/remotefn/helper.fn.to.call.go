package remotefn

import (
	"errors"
	"reflect"
)

func FnToCall(fn any, raw_args interface{}) (Call, error) {

	var zero Call

	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		return zero, errors.New("provided value is not a function")
	}
	args := make([]Param, fnType.NumIn())
	for i := 0; i < fnType.NumIn(); i++ {
		args[i] = Param{
			Name:      fnType.In(i).Name(),
			ParamType: fnType.In(i).String(),
		}
	}

	return zero, nil
}
