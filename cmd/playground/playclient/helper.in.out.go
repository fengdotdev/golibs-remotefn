package playclient

import (
	"errors"
	"fmt"
	"log"

	"github.com/fengdotdev/golibs-remotefn/sandbox/draf2/remote"
)

func NewDataExec() *DataExec {
	return &DataExec{}
}

type DataExec struct{}

func (d *DataExec) DataInOut(key string, data []byte) ([]byte, error) {

	if key == "Foo" {
		log.Println("Processing data for key: Foo")
		// Simulate processing the data for the "Foo" key

		m := remote.ResultSingle(5)
		outdata, err := remote.MapToDataOnJson(m)
		if err != nil {
			panic("this should never happen: " + err.Error())
		}
		fmt.Printf("Processed data for key: %s, result: %s\n", key, string(outdata))
		return outdata, nil

	}
	return nil, errors.New("unknown key: " + key)
}
