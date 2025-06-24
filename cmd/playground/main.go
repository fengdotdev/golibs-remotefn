package main

import (
	"fmt"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draft1/remotefn"
)

func main() {
	registry := remotefn.NewRegistryFn()

	// Define a contract for a function like: func add(a int, b int) int
	contract := remotefn.Contract{
		ConFnName: "add",
		ConParams: remotefn.Params{
			Params: []remotefn.Param{
				{Name: "a", ParamType: "int"},
				{Name: "b", ParamType: "int"},
			},
		},
		ConReplyParams: remotefn.ReplyParams{
			Params: []remotefn.Param{
				{Name: "result", ParamType: "int"},
			},
		},
	}

	// Register a function
	addFn := func(a, b int) int { return a + b }
	err := registry.Register(contract, addFn)
	if err != nil {
		fmt.Println("Register error:", err)
		return
	}

	// Call the function
	call := remotefn.Call{
		CallFnName: "add",
		CallArgs: remotefn.Args{
			ArgsArray: []remotefn.Arg{
				{ArgName: "a", ArgValue: 1},
				{ArgName: "b", ArgValue: 2},
			},
		},
	}

	reply, err := registry.CallFn(call)
	if err != nil {
		fmt.Println("Call error:", err)
		return
	}

	fmt.Println("Result:", reply.Args[0].Value) // Output: Result: 3
}
