package twelve

import (
	"fmt"
	"strings"
)

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

var verse = "On the %s day of Christmas my true love gave to me: %s."

func Verse(i int) string {
	if i == 1 {
		return fmt.Sprintf(verse, days[0], gifts[0])
	}

	verseGifts := make([]string, i)

	for j := i; j > 1; j-- {
        verseGifts[i-j] = gifts[j-1]
	}

    verseGifts[i-1] = fmt.Sprintf("and %s", gifts[0])

	return fmt.Sprintf(verse, days[i-1], strings.Join(verseGifts, ", "))
}

func Song() string {
	out := make([]string, 12)
	for i := 0; i < 12; i++ {
        out[i] = Verse(i+1)
	}

	return strings.Join(out, "\n")
}
