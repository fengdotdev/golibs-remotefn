package remotefn

import (
	"errors"
	"log"
	"reflect"
)

// given a Func add(a int, b int) =>  add(1,2) =>   paramA = {Name: A, ParamType: int} paramb = {Name:b, ParamType: int}
type Param struct {
	Name      string
	ParamType string
}

type Params struct {
	Params []Param
}

// value to a func ex:  given add(a int, b int) =>  add(1,2)  =>  argA = {Name: a, Value:1} argB = {Name: b, Value:2}
type Arg struct {
	Name  string
	Value interface{}
}

type Args struct {
	Args []Arg
}

type ReplyParams struct {
	Params []Param
}

type ContractDTO struct {
	FunctionName string
	Params       Params
	ReplyParams  ReplyParams
}

type CallDTO struct {
	FunctionName string
	Args         Args
}

type ReplyDTO struct {
	Args []Arg
}

type RegistryFn struct {
	functions map[string]interface{}
	contracts map[string]ContractDTO
}

func NewRegistryFn() *RegistryFn {
	return &RegistryFn{
		functions: make(map[string]interface{}),
		contracts: make(map[string]ContractDTO),
	}
}

func (r *RegistryFn) Register(contract ContractDTO, fn any) error {

	fnName := contract.FunctionName

	// register only non-existent
	_, ok := r.functions[fnName]
	if ok {
		return errors.New("fn already exist")
	}

	r.functions[fnName] = fn
	r.contracts[fnName] = contract
	return nil
}

func (r *RegistryFn) CallFn(call CallDTO) (ReplyDTO, error) {
	var zero ReplyDTO

	// Check if the function exists
	fnName := call.FunctionName
	fnAny, exists := r.functions[fnName]
	if !exists {
		return zero, errors.New("function does not exist")
	}

	// Get the contract
	contract, exists := r.contracts[fnName]
	if !exists {
		return zero, errors.New("contract does not exist")
	}

	// Validate arguments
	if len(call.Args.Args) != len(contract.Params.Params) {
		return zero, errors.New("argument count mismatch")
	}

	// Use reflection to get the function
	fnVal := reflect.ValueOf(fnAny)
	if fnVal.Kind() != reflect.Func {
		return zero, errors.New("registered value is not a function")
	}

	// Prepare input arguments
	args := make([]reflect.Value, len(call.Args.Args))
	for i, arg := range call.Args.Args {
		// Validate argument name and type (simplified; you may need more robust type checking)
		if i < len(contract.Params.Params) && arg.Name != contract.Params.Params[i].Name {
			return zero, errors.New("argument name mismatch")
		}
		log.Println("value  =>> ", arg.Value)
		args[i] = reflect.ValueOf(arg.Value)
	}

	// Call the function
	results := fnVal.Call(args)

	// Prepare reply
	reply := ReplyDTO{Args: make([]Arg, len(contract.ReplyParams.Params))}
	for i, result := range results {
		if i >= len(contract.ReplyParams.Params) {
			return zero, errors.New("unexpected number of return values")
		}

		log.Println(result)

		reply.Args[i] = Arg{
			Name:  contract.ReplyParams.Params[i].Name,
			Value: result.Interface(),
		}
	}

	return reply, nil
}
