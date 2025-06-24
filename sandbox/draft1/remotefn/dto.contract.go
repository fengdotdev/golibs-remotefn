package remotefn

type Contract struct {
	ConFnName      string      `json:"con_fn_name"`
	ConParams      Params      `json:"con_params"`
	ConReplyParams ReplyParams `json:"con_reply_params"`
}
