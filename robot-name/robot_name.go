package robotname

import (
	"errors"
	"fmt"
)

var (
	letter1 = 'A'
	letter2 = 'A'
	num     = 0
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := genName()
		if err != nil {
			return "", err
		}
		r.name = name
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name, _ = genName()
}

func genName() (string, error) {
	if letter1 > 'Z' {
		return "", errors.New("no more names available")
	}

	defer nextName()

	return fmt.Sprintf("%c%c%03d", letter1, letter2, num), nil
}

func nextName() {
	num++
	if num > 999 {
		num = 0
		letter2++
	}
	if letter2 > 'Z' {
		letter2 = 'A'
		letter1++
	}
}
