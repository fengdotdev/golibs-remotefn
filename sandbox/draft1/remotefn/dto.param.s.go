package remotefn

// given a Func add(a int, b int) =>  add(1,2) =>   paramA = {Name: A, ParamType: int} paramb = {Name:b, ParamType: int}
type Param struct {
	Name      string `json:"name"`
	ParamType string `json:"param_type"`
}

type Params struct {
	Params []Param `json:"params"`
}
