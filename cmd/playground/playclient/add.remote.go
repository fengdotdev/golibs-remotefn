package playclient

func AddRemote(a, b int) (int, error) {
	m := map[string]interface{}{
		"a": a,
		"b": b,
	}
	value, err := PipeSingleResult[int]("Add", m)
	if err != nil {
		return 0, err
	}

	return value, nil
}
