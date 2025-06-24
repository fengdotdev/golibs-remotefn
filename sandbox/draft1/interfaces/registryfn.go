package interfaces

type RegistryFn interface {
	Register(contract Contract, fn any) error
	CallFn(call Call) (Reply, error)
}
