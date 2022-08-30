package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2427"
  title: first-letter-to-appear-twice
  lang: golang
  type: large
  inputString: |-
    "abccbaacz"
    "abcdd"
    "ccaa"
    "abcdefghijklmnopqrstuvwxyzz"
*/

func repeatedCharacter(s string) byte {
	m := make([]bool, 26)
	for i := range s {
		if m[s[i]-'a'] {
			return s[i]
		}
		m[s[i]-'a'] = true
	}
	return 0
}

func Test_2427(t *testing.T) {
	NewTestcases(t).
		Add(byte('c'), "abccbaacz").
		Add(byte('d'), "abcdd").
		Add(byte('c'), "ccaa").
		Add(byte('z'), "abcdefghijklmnopqrstuvwxyzz").
		Each(func(a *assert.Assertions, td TestData) {
			actual := repeatedCharacter(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
