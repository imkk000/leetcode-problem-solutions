package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "344"
  title: reverse-string
  lang: golang
  type: large
  inputString: |-
    ["h","e","l","l","o"]
    ["H","a","n","n","a","h"]
*/

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

func Test_344(t *testing.T) {
	NewTestcases(t).
		Add("olleh", "hello").
		Add("hannaH", "Hannah").
		Add("a", "a").
		Add("ba", "ab").
		Each(func(a *assert.Assertions, td TestData) {
			actual := []byte(td.Input.(string))
			reverseString(actual)

			a.Equal([]byte(td.Expectation.(string)), actual)
		})
}
