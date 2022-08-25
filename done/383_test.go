package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "383"
  title: ransom-note
  lang: golang
  type: large
  inputString: |-
    "a"
    "b"
    "aa"
    "ab"
    "aa"
    "aab"
    "a"
    "a"
    "a"
    "aa"
    "aa"
    "a"
*/

func canConstruct(ransomNote string, magazine string) bool {
	dic := make([]int, 26)
	for _, r := range magazine {
		dic[r-'a']++
	}
	for _, r := range ransomNote {
		if dic[r-'a'] <= 0 {
			return false
		}
		dic[r-'a']--
	}
	return true
}

func Test_383(t *testing.T) {
	NewTestcases(t).
		Add(false, []string{"a", "b"}).
		Add(false, []string{"aa", "ab"}).
		Add(true, []string{"aa", "aab"}).
		Add(false, []string{"aa", "a"}).
		Add(true, []string{"a", "aa"}).
		Add(true, []string{"a", "a"}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.([]string)
			randomNote, magazine := input[0], input[1]
			actual := canConstruct(randomNote, magazine)

			a.Equal(td.Expectation, actual)
		})
}
