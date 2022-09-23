package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "557"
  title: reverse-words-in-a-string-iii
  lang: golang
  type: large
  inputString: |-
    "Let's take LeetCode contest"
    "God Ding"
*/

func reverseWords(in string) string {
	s := []byte(in)
	l := len(s)
	var start, end int
	for end = 0; end < l; end++ {
		if in[end] != ' ' {
			continue
		}
		reverse(s[start:end])
		start = end + 1
	}
	if start < end {
		reverse(s[start:end])
	}
	return string(s)
}

func reverse(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
}

func Test_557(t *testing.T) {
	NewTestcases(t).
		Add("s'teL ekat edoCteeL tsetnoc", "Let's take LeetCode contest").
		Add("doG gniD", "God Ding").
		Add("a b c d", "a b c d").
		Add("a", "a").
		Each(func(a *assert.Assertions, td TestData) {
			actual := reverseWords(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
