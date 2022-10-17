package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1960"
  title: check-if-the-sentence-is-pangram
  lang: golang
  type: large
  inputString: |-
    "thequickbrownfoxjumpsoverthelazydog"
    "leetcode"
*/

func checkIfPangram(sentence string) bool {
	var c int
	var v [26]bool
	for _, r := range sentence {
		if v[r-'a'] {
			continue
		}
		v[r-'a'] = true
		c++
	}
	return c == 26
}

func Test_1960(t *testing.T) {
	NewTestcases(t).
		Add(true, "thequickbrownfoxjumpsoverthelazydog").
		Add(false, "leetcode").
		Each(func(a *assert.Assertions, td TestData) {
			actual := checkIfPangram(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
