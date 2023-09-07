package diamond

import (
	"bytes"
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

	lineWidth := 1 + 2*(char-'A')
	out := make([]string, lineWidth)

	for c := byte('A'); c <= char; c++ {
		row := bytes.Repeat([]byte{' '}, int(lineWidth))
		colNum := char - c

		row[colNum] = c
		row[lineWidth-colNum-1] = c

		rowNum := c - 'A'
		out[rowNum] = string(row)
		out[lineWidth-rowNum-1] = string(row)
	}

	return strings.Join(out, "\n"), nil
}
