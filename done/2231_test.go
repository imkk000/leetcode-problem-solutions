package leetcode_test

import (
	"testing"

	_ "embed"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2231"
  title: find-first-palindromic-string-in-the-array
  lang: golang
  type: large
  inputString: |-
    ["abc","car","ada","racecar","cool"]
    ["notapalindrome","racecar"]
    ["def","ghi"]
*/

func firstPalindrome(words []string) string {
	for _, w := range words {
		if isPalindrome2231(w) {
			return w
		}
	}
	return ""
}

func isPalindrome2231(w string) bool {
	l := len(w)
	for i := 0; i < l/2; i++ {
		if w[i] != w[l-i-1] {
			return false
		}
	}
	return true
}

// go:exclude
var (
	//go:embed testdata/2231_1_in
	rawBenchInput string
	benchInput    = MakeStringSlice(rawBenchInput)
)

func Test_2231(t *testing.T) {
	NewTestcases(t).
		Add("ada", MakeStringSlice(`["ada"]`)).
		Add("ada", MakeStringSlice(`["abc","car","ada","racecar","cool"]`)).
		Add("racecar", MakeStringSlice(`["notapalindrome","racecar"]`)).
		Add("", MakeStringSlice(`["def","ghi"]`)).
		Add("", benchInput).
		Each(func(a *assert.Assertions, td TestData) {
			actual := firstPalindrome(td.Input.([]string))

			a.Equal(td.Expectation, actual)
		})
}
