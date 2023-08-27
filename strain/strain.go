package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	if i == nil {
		return i
	}

	o := make(Ints, 0)

	for _, v := range i {
		if filter(v) {
			o = append(o, v)
		}
	}

	return o
}

func (i Ints) Discard(filter func(int) bool) Ints {
	if i == nil {
		return i
	}

	o := make(Ints, 0)

	for _, v := range i {
		if !filter(v) {
			o = append(o, v)
		}
	}

	return o
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	if l == nil {
		return l
	}

	o := make(Lists, 0)

	for _, v := range l {
		if filter(v) {
			o = append(o, v)
		}
	}

	return o

}
func (s Strings) Keep(filter func(string) bool) Strings {
	if s == nil {
		return s
	}

	o := make(Strings, 0)

	for _, v := range s {
		if filter(v) {
			o = append(o, v)
		}
	}

	return o
}
