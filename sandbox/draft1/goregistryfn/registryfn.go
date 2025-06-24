package goregistryfn

type RegistryFn struct {
	functions map[string]interface{}
	contracts map[string]Contract
}
