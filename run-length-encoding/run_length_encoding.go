package encode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}

	counter := 0
	var previous rune = rune(input[0])

	s := strings.Builder{}
	for _, r := range input {
		if r != previous {
			if counter == 1 {
				s.WriteRune(previous)
			} else {
				s.WriteString(fmt.Sprintf("%d", counter))
				s.WriteRune(previous)
			}

			counter = 0
			previous = r
		}

		counter++
	}

	if counter == 1 {
		s.WriteRune(previous)
	} else {
		s.WriteString(fmt.Sprintf("%d", counter))
		s.WriteRune(previous)
	}

	return s.String()
}

func RunLengthDecode(input string) string {
    if input == "" {
        return ""
    }

    decoder := regexp.MustCompile(`(\d*)([\sa-zA-Z])`)
    found := decoder.FindAllStringSubmatch(input, -1)

    s := strings.Builder{}
    for _, group := range found {
        if group[1] == "" {
            s.WriteString(group[2])
            continue
        }
        n, _ := strconv.Atoi(group[1])
        s.WriteString(strings.Repeat(group[2], n))
    }

    return s.String()
}
