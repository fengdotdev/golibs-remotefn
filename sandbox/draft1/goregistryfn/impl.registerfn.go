package goregistryfn

import (
	"errors"
	"log"
	"reflect"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draft1/interfaces"
)

var _ interfaces.RegistryFn = NewRegistryFn()

// CallFn implements interfaces.RegistryFn.
func (r *RegistryFn) CallFn(call Call) (Reply, error) {
	var zero Reply

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
	reply := Reply{Args: make([]Arg, len(contract.ReplyParams.Params))}
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

// Register implements interfaces.RegistryFn.
func (r *RegistryFn) Register(contract Contract, fn any) error {
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
