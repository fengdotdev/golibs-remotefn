package remote

import (
	"encoding/json"
)

func MapToDataOnJson(m map[string]interface{}) ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DataToMapOnJson(data []byte) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
