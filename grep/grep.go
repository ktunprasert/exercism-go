package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var programFlags = map[string]bool{
	"-n": false, // Prepend line number
	"-l": false, // Output only the name of the files
	"-i": false, // Insensitive matching
	"-v": false, // Invert the program, collect non-matching files
	"-x": false, // Search when the string matches entire line
}

func Search(pattern string, flags, files []string) []string {
	defer resetFlags()
	setFlags(flags)

	result := make([]string, 0)
	scanned := make(map[string]struct{})

	for _, filename := range files {
		if _, ok := scanned[filename]; ok {
			continue
		}

		file, err := os.Open(filename)
		if err != nil {
			return nil
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		if programFlags["-i"] {
			pattern = fmt.Sprintf("%s%s", `(?i)`, pattern)
		}

		if programFlags["-x"] {
			pattern = fmt.Sprintf(`^%s$`, pattern)
		}

		re := regexp.MustCompile(pattern)

		lineN := 1

		var fileMatched bool
		for scanner.Scan() {
			line := scanner.Text()

			// XOR between LINE_MATCHED and INVERT_PROGRAM
			if re.MatchString(line) != programFlags["-v"] {
				switch {
				case programFlags["-l"]:
					if fileMatched {
						break
					}

					result = append(result, filename)
				case programFlags["-n"]:
					if len(files) > 1 {
						line = fmt.Sprintf("%s:%d:%s", filename, lineN, line)
					} else {
						line = fmt.Sprintf("%d:%s", lineN, line)
					}
					result = append(result, line)
				default:
					if len(files) > 1 {
						line = fmt.Sprintf("%s:%s", filename, line)
					}
					result = append(result, line)
				}

				fileMatched = true
			}

			lineN++
		}

		scanned[filename] = struct{}{}
	}

	return result
}

func setFlags(flags []string) {
	for _, f := range flags {
		if _, ok := programFlags[f]; ok {
			programFlags[f] = true
		}
	}
}

func resetFlags() {
	for key := range programFlags {
		programFlags[key] = false
	}
}
