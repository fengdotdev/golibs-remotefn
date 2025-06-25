package main

import (
	"fmt"

	"github.com/fengdotdev/golibs-remotefn/cmd/playground/playclient"
)

func main() {

	result, err := playclient.FooRemote(1, 2)
	if err != nil {
		panic(fmt.Sprintf("Error in FooRemote: %v", err))
	}
	println("Result:", result)

}
