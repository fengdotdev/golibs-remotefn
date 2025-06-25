package dataexec

import (
	"context"
	"fmt"
	"log"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"
)

var _ DataExec = (*GoDataExec)(nil)

func (d *GoDataExec) AddInMiddleware(fn DataInOutFn) {
	d.inMiddleware = append(d.inMiddleware, fn)
}
func (d *GoDataExec) AddOutMiddleware(fn DataInOutFn) {
	d.outMiddleware = append(d.outMiddleware, fn)
}

func (d *GoDataExec) DataIn(key string, data []byte) ([]byte, error) {
	// Execute all input middleware functions

	if len(d.inMiddleware) == 0 {
		return data, nil // No input middleware, return data as is
	}

	for _, fn := range d.inMiddleware {
		var err error
		data, err = fn(key, data)
		if err != nil {
			return nil, fmt.Errorf("error in input middleware for key %s: %w", key, err)
		}
	}
	return d.DataInOut(key, data)
}

func (d *GoDataExec) DataOut(key string, data []byte) ([]byte, error) {
	// Execute all output middleware functions

	if len(d.outMiddleware) == 0 {
		return data, nil // No output middleware, return data as is
	}

	for _, fn := range d.outMiddleware {
		var err error
		data, err = fn(key, data)
		if err != nil {
			return nil, fmt.Errorf("error in output middleware for key %s: %w", key, err)
		}
	}
	return data, nil
}

func (d *GoDataExec) DataInOut(key string, dataIn []byte) ([]byte, error) {

	processedDataIn, err := d.DataIn(key, dataIn)
	if err != nil {
		return nil, fmt.Errorf("error in DataIn for key %s: %w", key, err)
	}

	fn, err := d.registry.GetRemoteFn(key)
	if err != err {
		return nil, fmt.Errorf("failed to get remote function for key %s: %w", key, err)
	}

	if fn == nil {
		return nil, fmt.Errorf("remote function for key %s is nil", key)
	}

	log.Printf("Executing remote function for key: %s", key)

	m, err := remote.DataToMapOnJson(processedDataIn)
	if err != nil {
		return nil, fmt.Errorf("error converting data to map for key %s: %w", key, err)
	}

	ctx := context.Background()

	result, err := fn(ctx, m)
	if err != nil {
		return nil, fmt.Errorf("error executing remote function for key %s: %w", key, err)
	}
	if result == nil {
		return nil, fmt.Errorf("remote function for key %s returned nil result", key)
	}

	dataOut, err := remote.MapToDataOnJson(result)
	if err != nil {
		return nil, fmt.Errorf("error converting result to data for key %s: %w", key, err)
	}

	processedDataOut, err := d.DataOut(key, dataOut)
	if err != nil {
		return nil, fmt.Errorf("error in DataOut for key %s: %w", key, err)
	}

	log.Printf("Successfully executed remote function for key: %s", key)
	return processedDataOut, nil
}
