package stringset

import (
	"fmt"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	set map[string]struct{}
}

func New() Set {
	return Set{make(map[string]struct{})}
}

func NewFromSlice(l []string) Set {
	set := make(map[string]struct{})
	for _, element := range l {
		set[element] = struct{}{}
	}

	return Set{set}
}

func (s Set) String() string {
	if len(s.set) == 0 {
		return "{}"
	}

	str := strings.Builder{}
	str.WriteString("{")
	i := 0
	for k := range s.set {
		str.WriteString(fmt.Sprintf("%q", k))
		if i < len(s.set)-1 {
			str.WriteString(", ")
		}
		i++
	}
	str.WriteString("}")

	return str.String()
}

func (s Set) IsEmpty() bool {
	return len(s.set) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s.set[elem]

	return ok
}

func (s Set) Add(elem string) {
	s.set[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	for elem := range s1.set {
		if !s2.Has(elem) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
    return Intersection(s1, s2).IsEmpty()
}

func Equal(s1, s2 Set) bool {
	if len(s1.set) != len(s2.set) {
		return false
	}

	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	set := make(map[string]struct{})
	for elem := range s1.set {
		if s2.Has(elem) {
			set[elem] = struct{}{}
		}
	}

	for elem := range s2.set {
		if s1.Has(elem) {
			set[elem] = struct{}{}

		}
	}

	return Set{set}
}

func Difference(s1, s2 Set) Set {
	set := make(map[string]struct{})
	for elem := range s1.set {
		set[elem] = struct{}{}
	}

	for elem := range s2.set {
		delete(set, elem)
	}

	return Set{set}
}

func Union(s1, s2 Set) Set {
	for elem := range s2.set {
		s1.Add(elem)
	}

	return s1
}
