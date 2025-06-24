package remotefn

import (
	"errors"
	"reflect"
)

func FnToCrontract(fn any, fnName string) (Contract, error) {

	var zero Contract

	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		return zero, errors.New("provided value is not a function")
	}

	params := make([]Param, fnType.NumIn())
	for i := 0; i < fnType.NumIn(); i++ {
		params[i] = Param{
			Name:      fnType.In(i).Name(),
			ParamType: fnType.In(i).String(),
		}
	}

	replyParams := make([]Param, fnType.NumOut())
	for i := 0; i < fnType.NumOut(); i++ {
		replyParams[i] = Param{
			Name:      fnType.Out(i).Name(),
			ParamType: fnType.Out(i).String(),
		}
	}

	return Contract{
		FunctionName: fnName,
		Params:       Params{Params: params},
		ReplyParams:  ReplyParams{Params: replyParams},
	}, nil
}
