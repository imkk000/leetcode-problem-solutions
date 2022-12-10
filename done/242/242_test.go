package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "242"
  title: valid-anagram
  lang: golang
  type: large
  inputString: |-
    "anagram"
    "nagaram"
    "rat"
    "car"
    "ac"
    "a"
*/

func isAnagram(s string, t string) bool {
	var c [26]int
	for _, r := range s {
		c[r-'a']++
	}
	for _, r := range t {
		c[r-'a']--
		if c[r-'a'] < 0 {
			return false
		}
	}
	for i := range c {
		if c[i] > 0 {
			return false
		}
	}
	return true
}

func Test_242(t *testing.T) {
	type Data struct {
		s, t string
	}
	NewTestcases(t).
		Add(true, Data{
			s: "anagram",
			t: "nagaram",
		}).
		Add(false, Data{
			s: "rat",
			t: "car",
		}).
		Add(false, Data{
			s: "ac",
			t: "a",
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			s, t := input.s, input.t
			actual := isAnagram(s, t)

			a.Equal(td.Expectation, actual)
		})
}
