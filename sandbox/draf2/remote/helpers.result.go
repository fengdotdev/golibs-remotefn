package remote

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
