package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1823"
  title: determine-if-string-halves-are-alike
  lang: golang
  type: large
  inputString: |-
    "book"
    "textbook"
*/

func halvesAreAlike(s string) bool {
	l := len(s)
	var left, right int
	for i := 0; i < l/2; i++ {
		left += isVowel(s[i])
		right += isVowel(s[(l/2)+i])
	}
	return left == right
}

func isVowel(r byte) int {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return 1
	}
	return 0
}

func Test_1823(t *testing.T) {
	NewTestcases(t).
		Add(true, "book").
		Add(false, "textbook").
		Add(true, "ao").
		Add(true, "tk").
		Add(false, "ur").
		Each(func(a *assert.Assertions, td TestData) {
			actual := halvesAreAlike(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
