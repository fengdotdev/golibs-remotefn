package remotefn

// value to a func ex:  given add(a int, b int) =>  add(1,2)  =>  argA = {Name: a, Value:1} argB = {Name: b, Value:2}
type Arg struct {
	ArgName  string      `json:"arg_name"`
	ArgValue interface{} `json:"arg_value"`
}

type Args struct {
	ArgsArray []Arg `json:"args_array"`
}
