package main

import (
	"fmt"

	"github.com/fengdotdev/golibs-remotefn/cmd/playground/playclient"
)

func main() {

	result, err := playclient.AddRemote(1, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		println("Result:", result)

	}

}
