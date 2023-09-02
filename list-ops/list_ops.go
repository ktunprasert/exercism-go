package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	folded := initial
	for i := 0; i < len(s); i++ {
		folded = fn(folded, s[i])
	}

	return folded
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	folded := initial
	for i := len(s) - 1; i >= 0; i-- {
		folded = fn(s[i], folded)
	}

	return folded
}

func (s IntList) Filter(fn func(int) bool) IntList {
	filtered := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
            filtered.Append(IntList{v})
		}

	}
	return filtered
}

func (s IntList) Length() int {
	return len(s)
}

func (s *IntList) Map(fn func(int) int) IntList {
	for i, v := range *s {
		(*s)[i] = fn(v)
	}

	return *s
}

func (s IntList) Reverse() IntList {
	reversed := make(IntList, len(s))
	for i, v := range s {
		reversed[len(s)-i-1] = v
	}

	return reversed
}

func (s IntList) Append(lst IntList) IntList {
	appended := make(IntList, len(s)+len(lst))
	copy(appended, s)
	copy(appended[len(s):], lst)

	return appended
}

func (s IntList) Concat(lists []IntList) IntList {
	concatenated := make(IntList, 0)
	concatenated = append(concatenated, s...)

	for _, l := range lists {
		concatenated = append(concatenated, l...)
	}

	return concatenated
}
