package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if rune(char) > 'Z' || rune(char) < 'A' {
		return "", errors.New("char out of range")
	}

	if rune(char) == 'A' {
		return "A", nil
	}

	lineWidth := 1 + 2*(rune(char)-'A')
	s := strings.Builder{}

	for i := 'A'; i <= rune(char); i++ {
		middle := renderChar(i)
		spaces := (int(lineWidth) - len(middle)) / 2
		s.WriteString(strings.Repeat(" ", int(spaces)))
		s.WriteString(renderChar(i))
		s.WriteString(strings.Repeat(" ", int(spaces)))
		s.WriteString("\n")
	}

	for i := rune(char) - 1; i >= 'A'; i-- {
		middle := renderChar(i)
		spaces := (int(lineWidth) - len(middle)) / 2
		s.WriteString(strings.Repeat(" ", int(spaces)))
		s.WriteString(renderChar(i))
		s.WriteString(strings.Repeat(" ", int(spaces)))
		s.WriteString("\n")
	}

	return s.String()[:len(s.String())-1], nil
}

func renderChar(c rune) string {
	if c == 'A' {
		return "A"
	}

	dist := c - 'A'
	return string(c) + strings.Repeat(" ", 2*(int(dist)-1)+1) + string(c)
}
