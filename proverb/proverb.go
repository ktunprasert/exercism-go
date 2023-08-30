// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	rhymes := make([]string, len(rhyme))

	for i := 0; i < len(rhyme)-1; i++ {
		rhymes[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
	}

	rhymes[len(rhyme)-1] = "And all for the want of a " + rhyme[0] + "."

	return rhymes
}
