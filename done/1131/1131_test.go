package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1131"
  title: count-substrings-with-only-one-distinct-letter
  lang: golang
  type: large
  inputString: |-
    "aaaba"
    "aaaaaaaaaa"
    "a"
    "ab"
    "aba"
    "abab"
    "aabb"
*/

func countLetters(s string) (v int) {
	s += "_"
	var start, end int
	l := len(s)
	for i := 0; i < l-1; i++ {
		if s[i] != s[i+1] {
			end = i + 1
			v += (end - start) * (end - start + 1) / 2
			start = end
		}
	}
	return
}

func Test_1131(t *testing.T) {
	NewTestcases(t).
		Add(1, "a").
		Add(2, "ab").
		Add(3, "aba").
		Add(4, "abab").
		Add(6, "aabb").
		Add(8, "aaaba").
		Add(55, "aaaaaaaaaa").
		Each(func(a *assert.Assertions, td TestData) {
			actual := countLetters(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
