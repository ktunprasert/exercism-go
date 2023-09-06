package flatten

func Flatten(nested interface{}) []interface{} {
	flat := make([]interface{}, 0)

	switch nested.(type) {
	case []interface{}:
		for _, v := range nested.([]interface{}) {
			flat = append(flat, Flatten(v)...)
		}

	case interface{}:
		flat = append(flat, nested)
	}

	return flat
}
