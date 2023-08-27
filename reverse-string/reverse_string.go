package reverse

func Reverse(input string) string {
	if len(input) == 0 {
		return ""
	}

	var out string
	for _, r := range input {
		out = string(r) + out
	}

	return out
}
