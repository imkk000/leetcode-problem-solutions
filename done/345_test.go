package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "345"
  title: reverse-vowels-of-a-string
  lang: golang
  type: large
  inputString: |-
    "hello"
    "leetcode"
    "ehi"
    "ehI"
    "ae"
    "a"
    "h"
*/

func reverseVowels(s string) string {
	const vowels = "aeiouAEIOU"
	isVowel := func(b byte) bool {
		for i := range vowels {
			if vowels[i] == b {
				return true
			}
		}
		return false
	}

	b := []byte(s)
	l, r := 0, len(b)-1
	for l < r {
		if !isVowel(b[l]) {
			l++
			continue
		}
		for !isVowel(b[r]) && l < r {
			r--
		}

		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return string(b)
}

func Test_345(t *testing.T) {
	NewTestcases(t).
		Add("holle", "hello").
		Add("leotcede", "leetcode").
		Add("a", "a").
		Add("h", "h").
		Add("ea", "ae").
		Add("ihe", "ehi").
		Add("Ihe", "ehI").
		Each(func(a *assert.Assertions, td TestData) {
			actual := reverseVowels(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
