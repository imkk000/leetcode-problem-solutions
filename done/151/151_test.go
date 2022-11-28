package leetcode_test

import (
	"regexp"
	"strings"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "151"
  title: reverse-words-in-a-string
  lang: golang
  type: large
  inputString: |-
    "the sky is blue"
    "  hello world  "
    "a good   example"
*/

func reverseWords(s string) string {
	re, _ := regexp.Compile(`\s+`)
	s = strings.TrimSpace(re.ReplaceAllString(s, " "))
	b := []byte(s)
	var start, end int
	for i := range b {
		end++
		if b[i] == ' ' {
			reverse(b[start : end-1])
			start = i + 1
			continue
		}
	}
	if start < end {
		reverse(b[start:end])
	}
	reverse(b)
	return string(b)
}

func reverse(b []byte) {
	l := len(b)
	for i := 0; i < l/2; i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
}

func Test_151(t *testing.T) {
	NewTestcases(t).
		Add("blue is sky the", "the sky is blue").
		Add("world hello", "  hello world  ").
		Add("example good a", "a good   example").
		Add("Alice Loves Bob", "  Bob    Loves  Alice   ").
		Add("a", "     a     ").
		Add("a", "a").
		Add("a", " a ").
		Each(func(a *assert.Assertions, td TestData) {
			actual := reverseWords(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
