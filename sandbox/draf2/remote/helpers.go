package remote

import (
	"context"
	"encoding/json"
)

func OrErr[T any](m map[string]interface{}, key string) (T, error) {
	var zero T

	if value, ok := m[key]; ok {
		return value.(T), nil
	}
	return zero, nil
}

func Or[T any](m map[string]interface{}, key string, defaultValue T) T {
	if value, ok := m[key]; ok {
		return value.(T)
	}
	return defaultValue
}

func ResultSingle[T any](value T) map[string]interface{} {
	return map[string]interface{}{
		"result": value,
	}
}

func ResultAndErr[T any](value T, err error) map[string]interface{} {
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}
	return map[string]interface{}{
		"result": value,
	}
}

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

func MockWire(ctx context.Context, datachan chan []byte, key string, fn func(key string, data []byte) []byte) chan []byte {

	output := make(chan []byte)
	go func(key string) {
		defer close(output)

		for {
			select {
			case data := <-datachan:
				if data == nil {
					return
				}
				result := fn(key, data)
				if result == nil {
					return
				}

				output <- result
			case <-ctx.Done():
				return
			}
		}
	}(key)
	return output
}
