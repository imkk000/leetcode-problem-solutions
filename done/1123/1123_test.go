package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1123"
  title: single-row-keyboard
  lang: golang
  type: large
  inputString: |-
    "abcdefghijklmnopqrstuvwxyz"
    "cba"
    "pqrstuvwxyzabcdefghijklmno"
    "leetcode"
    "abcdefghijklmnopqrstuvwxyz"
    "aaa"
    "abcdefghijklmnopqrstuvwxyz"
    "z"
*/

func calculateTime(keyboard string, word string) (v int) {
	var s [26]int
	for i, k := range keyboard {
		s[k-'a'] = i
	}
	var c int
	for _, r := range word {
		v += abs(c - s[r-'a'])
		c = int(s[r-'a'])
	}
	return
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Test_1123(t *testing.T) {
	type Data struct {
		keyboard, word string
	}
	NewTestcases(t).
		Add(4, Data{
			keyboard: "abcdefghijklmnopqrstuvwxyz",
			word:     "cba",
		}).
		Add(73, Data{
			keyboard: "pqrstuvwxyzabcdefghijklmno",
			word:     "leetcode",
		}).
		Add(0, Data{
			keyboard: "abcdefghijklmnopqrstuvwxyz",
			word:     "aaaa",
		}).
		Add(25, Data{
			keyboard: "abcdefghijklmnopqrstuvwxyz",
			word:     "z",
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			keyboard, word := input.keyboard, input.word
			actual := calculateTime(keyboard, word)

			a.Equal(td.Expectation, actual)
		})
}
