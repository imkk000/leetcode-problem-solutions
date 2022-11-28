package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1490"
  title: generate-a-string-with-characters-that-have-odd-counts
  lang: golang
  type: large
  inputString: |-
    1
    2
    3
    4
    5
    6
    7
    500
*/

const s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

func generateTheString(n int) string {
	if n&1 == 1 {
		return s[:n]
	}
	return s[:n-1] + "b"
}

func Test_1490(t *testing.T) {
	NewTestcases(t).
		Add("a", 1).
		Add("ab", 2).
		Add("aaa", 3).
		Add("aaab", 4).
		Add("aaaaa", 5).
		Each(func(a *assert.Assertions, td TestData) {
			actual := generateTheString(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
