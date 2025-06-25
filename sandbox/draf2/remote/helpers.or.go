package remote

import (
	"errors"
	"fmt"
	"log"
)

func OrErr[T any](m map[string]interface{}, key string) (T, error) {
	var zero T

	valueRaw, ok := m[key]

	if !ok {

		return zero, errors.New("key not found: " + key)
	}

	value, ok := valueRaw.(T) // type assertion FOR  INT RETURN FLOAT64 ?
	if !ok {
		log.Printf("Expected type %T for key %s, but got %T", zero, key, valueRaw)
		return zero, errors.New("value for key " + key + " is not of type " + fmt.Sprintf("%T", zero))
	}
	return value, nil
}

func Or[T any](m map[string]interface{}, key string, defaultValue T) T {

	valueRaw, ok := m[key]
	if !ok {

		return defaultValue
	}

	value, ok := valueRaw.(T)
	if !ok {
		return defaultValue
	}

	return value
}
