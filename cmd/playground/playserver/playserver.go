package playserver

import (
	"github.com/fengdotdev/golibs-remotefn/cmd/playground/playserver/dataexec"
	"github.com/fengdotdev/golibs-remotefn/cmd/playground/playserver/handlers"
	"github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"
)

func NewDataExec() dataexec.DataExec {
	reg := remote.NewGoRegister()

	// register functions
	reg.RegisterRemoteFn("Add", handlers.AddHandler)

	return dataexec.NewDataExec(reg)
}
