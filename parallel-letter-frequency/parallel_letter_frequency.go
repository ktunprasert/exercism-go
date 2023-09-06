package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap)

	var wg sync.WaitGroup
	wg.Add(len(texts))

	for _, text := range texts {
		go func(text string) {
			defer wg.Done()
			c <- Frequency(text)
		}(text)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	freq := FreqMap{}
	for f := range c {
		for k, v := range f {
			freq[k] += v
		}
	}

	return freq
}
