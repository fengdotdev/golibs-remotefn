package dataexec

import (
	"github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"
)

type GoDataExec struct {
	inMiddleware  []DataInOutFn
	registry      remote.RegisterRemote
	outMiddleware []DataInOutFn
}
