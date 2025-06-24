package remote

import "fmt"

var _ RegisterRemote = (*GoRegister)(nil)

func NewGoRegister() *GoRegister {
	return &GoRegister{
		remoteFns: make(map[string]RemoteFn),
	}
}

type GoRegister struct {
	remoteFns map[string]RemoteFn
}

// GetRemoteFn implements RegisterRemote.
func (g *GoRegister) GetRemoteFn(name string) (RemoteFn, error) {
	if g.remoteFns == nil {
		return nil, fmt.Errorf("no remote functions registered")
	}
	fn, exists := g.remoteFns[name]
	if !exists {
		return nil, fmt.Errorf("remote function %s not found", name)
	}
	return fn, nil
}

// RegisterRemoteFn implements RegisterRemote.
func (g *GoRegister) RegisterRemoteFn(name string, fn RemoteFn) error {
	if g.remoteFns == nil {
		g.remoteFns = make(map[string]RemoteFn)
	}
	if _, exists := g.remoteFns[name]; exists {
		return fmt.Errorf("remote function %s already registered", name)
	}
	g.remoteFns[name] = fn
	return nil
}
