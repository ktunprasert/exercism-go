package series

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}

	result := make([]string, 0)
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:i+n])
	}

	return result
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
