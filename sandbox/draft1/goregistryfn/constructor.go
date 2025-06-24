package goregistryfn

func NewRegistryFn() *RegistryFn {
	return &RegistryFn{
		functions: make(map[string]interface{}),
		contracts: make(map[string]Contract),
	}
}
