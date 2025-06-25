package dataexec

import "github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"

func NewDataExec(reg remote.RegisterRemote) *GoDataExec {
	inMiddleware := make([]DataInOutFn, 0)
	outMiddleware := make([]DataInOutFn, 0)

	return &GoDataExec{
		inMiddleware:  inMiddleware,
		outMiddleware: outMiddleware,
		registry:      reg,
	}
}
