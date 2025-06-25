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
	// Handle common type conversions for basic types
	switch any(zero).(type) {
	case int:
		switch v := valueRaw.(type) {
		case int:
			return any(v).(T), nil
		case float64:
			return any(int(v)).(T), nil
		}
	case float64:
		switch v := valueRaw.(type) {
		case float64:
			return any(v).(T), nil
		case int:
			return any(float64(v)).(T), nil
		}
	case string:
		if v, ok := valueRaw.(string); ok {
			return any(v).(T), nil
		}
	}
	// Fallback to direct assertion
	value, ok := any(valueRaw).(T)
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
	value, ok := any(valueRaw).(T)
	if !ok {
		return defaultValue
	}

	return value
}
