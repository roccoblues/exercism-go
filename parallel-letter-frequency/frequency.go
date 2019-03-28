package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calls Frequency for each input string concurrently and
// returns the merged results.
func ConcurrentFrequency(inputs []string) FreqMap {
	cnt := len(inputs)
	results := make(chan FreqMap, cnt)

	for _, s := range inputs {
		go func(s string) {
			results <- Frequency(s)
		}(s)
	}

	m := FreqMap{}
	for i := 0; i < cnt; i++ {
		for s, n := range <-results {
			m[s] += n
		}
	}

	return m
}
