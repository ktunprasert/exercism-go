package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Search(pattern string, flags, files []string) []string {
	var flagLineNumber, flagOnlyName, flagInsensitive, flagInvert, flagEntireLine bool

	flagMultipleFiles := len(files) > 1

	for _, f := range flags {
		switch f {
		case "-n":
			flagLineNumber = true
		case "-l":
			flagOnlyName = true
		case "-i":
			flagInsensitive = true
		case "-v":
			flagInvert = true
		case "-x":
			flagEntireLine = true
		}
	}

	result := make([]string, 0)

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			return nil
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		if flagInsensitive {
			pattern = fmt.Sprintf("%s%s", `(?i)`, pattern)
		}

		if flagEntireLine {
			pattern = fmt.Sprintf(`^%s$`, pattern)
		}

		re := regexp.MustCompile(pattern)

		lineN := 1

		for scanner.Scan() {
			line := scanner.Text()

			// XOR between LINE_MATCHED and INVERT_PROGRAM
			if re.MatchString(line) != flagInvert {
				if flagOnlyName {
					result = append(result, filename)
					break
				}

				if flagLineNumber {
					line = fmt.Sprintf("%d:%s", lineN, line)
				}

				if flagMultipleFiles {
					line = fmt.Sprintf("%s:%s", filename, line)
				}

				result = append(result, line)
			}

			lineN++
		}
	}

	return result
}
