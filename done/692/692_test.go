package leetcode_test

import (
	"sort"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "692"
  title: top-k-frequent-words
  lang: golang
  type: large
  inputString: |-
    ["i","love","leetcode","i","love","coding"]
    2
    ["the","day","is","sunny","the","the","the","sunny","is","is"]
    4
    ["a","b"]
    2
*/

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	var freq []string
	for _, w := range words {
		if m[w] == 0 {
			freq = append(freq, w)
		}
		m[w]++
	}
	sort.Slice(freq, func(i, j int) bool {
		a, b := freq[i], freq[j]
		if m[a] != m[b] {
			return m[a] > m[b]
		}
		return a < b
	})
	return freq[:k]
}

func Test_692(t *testing.T) {
	type Data struct {
		words []string
		k     int
	}
	NewTestcases(t).
		AddExpectation(MakeStringSlice(`["i","love"]`)).
		AddInput(Data{
			words: MakeStringSlice(`["i","love","leetcode","i","love","coding"]`),
			k:     2,
		}).
		AddExpectation(MakeStringSlice(`["the","is","sunny","day"]`)).
		AddInput(Data{
			words: MakeStringSlice(`["the","day","is","sunny","the","the","the","sunny","is","is"]`),
			k:     4,
		}).
		AddExpectation(MakeStringSlice(`["a","b"]`)).
		AddInput(Data{
			words: MakeStringSlice(`["a","b"]`),
			k:     2,
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			words, k := input.words, input.k
			actual := topKFrequent(words, k)

			a.Equal(td.Expectation, actual)
		})
}
