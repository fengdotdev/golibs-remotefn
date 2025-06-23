package remotefn

import (
	"errors"
	"reflect"
)

func FnToCrontract(fn any, fnName string) (ContractDTO, error) {

	var zero ContractDTO

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

	return ContractDTO{
		FunctionName: fnName,
		Params:       Params{Params: params},
		ReplyParams:  ReplyParams{Params: replyParams},
	}, nil
}

func FnToCall(fn any, raw_args interface{}) (CallDTO, error) {

	var zero CallDTO

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

	return CallDTO{
		FunctionName: fnType.Name(),
		Args:         Args{Args: args, RawArgs: raw_args},
	}, nil
}
