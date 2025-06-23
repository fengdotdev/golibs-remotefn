package remotefn

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
