package sublist

func Sublist(l1, l2 []int) Relation {
	switch {
	case len(l1) == 0 && len(l2) == 0:
		return RelationEqual
	case len(l1) == 0 && len(l2) > 0:
		return RelationSublist
	case len(l1) > 0 && len(l2) == 0:
		return RelationSuperlist
	}

	switch {
	case len(l1) == len(l2):
		if contains(l1, l2) {
			return RelationEqual
		}
	case len(l1) < len(l2):
		if contains(l2, l1) {
			return RelationSublist
		}
	case len(l1) > len(l2):
		if contains(l1, l2) {
			return RelationSuperlist
		}
	}

	return RelationUnequal
}

func contains(big, small []int) bool {
	for i := 0; i <= len(big)-len(small); i++ {
		window := big[i : i+len(small)]

		var eq int
		for i, r := range window {
			if r != small[i] {
				break
			}
			eq++
		}

		if eq == len(small) {
			return true
		}
	}

	return false
}
