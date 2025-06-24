package main

import (
	"errors"
	"fmt"
	"reflect"
)

func Foo(a, b int) int {
	return a + b
}

func main() {
	fnType := reflect.TypeOf(Foo)
	if fnType.Kind() != reflect.Func {
		panic(errors.New("provided value is not a function"))
	}

	fmt.Println("Function Name:", fnType.Name())
	fmt.Println("Number of Input Parameters:", fnType.NumIn())
	for i := 0; i < fnType.NumIn(); i++ {
		fmt.Println("Input Parameter:",)

		fmt.Println( fnType.In(i).Field())

		paramType := fnType.In(i)
		fmt.Printf("Parameter %d: %s (%s)\n", i+1, paramType.Name(), paramType.String())
	}
}
