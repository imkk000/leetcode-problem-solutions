package leetcode_test

import (
	_ "embed"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1089"
  title: remove-vowels-from-a-string
  lang: golang
  type: large
  inputString: |-
    "leetcodeisacommunityforcoders"
    "aeiou"
*/

func removeVowels(s string) string {
	var c int
	v := make([]rune, len(s))
	for _, r := range s {
		if isVowels(r) {
			continue
		}
		v[c] = r
		c++
	}
	return string(v[:c])
}

func isVowels(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func Test_1089(t *testing.T) {
	NewTestcases(t).
		Add("ltcdscmmntyfrcdrs", "leetcodeisacommunityforcoders").
		Add("", "aeiou").
		Add(expected, input).
		Each(func(a *assert.Assertions, td TestData) {
			actual := removeVowels(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}

// go:exclude
//
//go:embed 1089_1_in
var input string

// go:exclude
//
//go:embed 1089_1_out
var expected string
