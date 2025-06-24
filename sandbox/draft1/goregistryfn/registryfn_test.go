package goregistryfn_test

import (
	"errors"
	"testing"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draft1/goregistryfn"
	"github.com/fengdotdev/golibs-remotefn/sandbox/draft1/remotefn"
	"github.com/fengdotdev/golibs-testing/assert"
)

func add(a, b int) int {
	return a + b
}

func somecall(call string) ([]byte, error) {
	if call == "fail" {
		return nil, errors.New("expected fail")
	}

	msg := call

	return []byte(msg), nil
}

func TestRegistryFn1(t *testing.T) {
	reg := goregistryfn.NewRegistryFn()

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
	addFn := add
	err := reg.Register(contract, addFn)
	assert.NoError(t, err)

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

	reply, err := reg.CallFn(call)
	assert.NoError(t, err)

	value := reply.Args[0].ArgValue
	assert.Equal(t, value.(int), 3)

}
func TestRegistryFn2(t *testing.T) {
	reg := goregistryfn.NewRegistryFn()

	contract := remotefn.Contract{
		ConFnName: "somecall",
		ConParams: remotefn.Params{
			Params: []remotefn.Param{
				{Name: "call", ParamType: "string"},
			},
		},
		ConReplyParams: remotefn.ReplyParams{
			Params: []remotefn.Param{
				{Name: "error", ParamType: "error"},
				{Name: "slice", ParamType: "[]byte"},
			},
		},
	}

	// Register a function
	fn := somecall
	err := reg.Register(contract, fn)
	assert.NoError(t, err)
	inMsg := "foo"

	// Call the function
	call := remotefn.Call{
		CallFnName: "somecall",
		CallArgs: remotefn.Args{
			ArgsArray: []remotefn.Arg{
				{ArgName: "call", ArgValue: inMsg},
			},
		},
	}

	reply, err := reg.CallFn(call)
	assert.NoError(t, err)

	err, _ = reply.Args[1].ArgValue.(error)

	assert.NoError(t, err)

	value, _ := reply.Args[0].ArgValue.([]byte)

	str := string(value)

	assert.Equal(t, str, inMsg)

}
