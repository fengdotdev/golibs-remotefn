package remotefn_test

import (
	"fmt"
	"testing"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draft1/remotefn"
	"github.com/fengdotdev/golibs-testing/assert"
)

func Foo(a, b int) int {
	return a + b
}

func Bar(a, b int) (int, error) {
	return a + b, nil
}

func PrintBar() {
	fmt.Println("This is a print function for Bar")
}

type BazFoo struct {
	A string `json:"a"`
	X int    `json:"x"`
	Y bool   `json:"y"`
}

func Baz(a string, b map[string]string, c BazFoo) (int, error, string) {

	lenstrAC := len(c.A) + len(a)

	mymap := make(map[string]string)

	if b == nil {
		return 0, fmt.Errorf("b cannot be nil"), ""
	}

	mymap["A"] = c.A
	mymap["b"] = fmt.Sprintf("%v", c.X)
	mymap["c"] = fmt.Sprintf("%v", c.Y)

	return lenstrAC, nil, fmt.Sprintf("%v", mymap)
}

func TestFnToContract(t *testing.T) {

	contract, err := remotefn.FnToCrontract(Foo, "Foo")
	assert.NoError(t, err)
	t.Logf("Contract: %+v", contract)

	assert.Equal(t, "Foo", contract.ConFnName)
	assert.Equal(t, 2, len(contract.ConParams.Params))
	//assert.Equal(t, "a", contract.ConParams.Params[0].Name)
	assert.Equal(t, "int", contract.ConParams.Params[0].ParamType)
	//assert.Equal(t, "b", contract.ConParams.Params[1].Name)
	assert.Equal(t, "int", contract.ConParams.Params[1].ParamType)
	assert.Equal(t, 1, len(contract.ConReplyParams.Params))
	assert.Equal(t, "int", contract.ConReplyParams.Params[0].ParamType)
	assert.Equal(t, "int", contract.ConReplyParams.Params[0].Name)
}
