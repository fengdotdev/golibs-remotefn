package remotefn

type Call struct {
	CallFnName string `json:"call_fn_name"`
	CallArgs   Args   `json:"call_args"`
}
