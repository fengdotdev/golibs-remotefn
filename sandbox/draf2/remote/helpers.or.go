package remote

import (
	"errors"
	"fmt"
	"log"
)

func OrErr[T any](m map[string]interface{}, key string) (T, error) {
	var zero T

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic while accessing key %s: %v", key, r)

		}

	}()

	valueRaw, ok := m[key]
	if !ok {

		return zero, errors.New("key not found: " + key)
	}

	value, err := Assert[T](valueRaw)
	if err != nil {
		log.Printf("Error asserting type for key %s: %v", key, err)
		return zero, fmt.Errorf("value for key %s is not of type %T: %w", key, zero, err)
	}

	return value, nil
}

func Or[T any](m map[string]interface{}, key string, defaultValue T) T {

	value, err := OrErr[T](m, key)
	if err != nil {
		log.Printf("Error retrieving key %s: %v, returning default value", key, err)
		return defaultValue
	}

	return value
}

func Assert[T any](valueRaw any) (T, error) {
	var zero T

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
	case bool:
		if v, ok := valueRaw.(bool); ok {
			return any(v).(T), nil
		}
	case []interface{}:
		if v, ok := valueRaw.([]interface{}); ok {
			return any(v).(T), nil
		}
	}
	// Fallback to direct assertion
	value, ok := any(valueRaw).(T)
	if !ok {
		return zero, fmt.Errorf("value %v is not of type %T", valueRaw, zero)
	}
	return value, nil
}
