package leetcode_test

import (
	"strings"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1781"
  title: check-if-two-string-arrays-are-equivalent
  lang: golang
  type: large
  inputString: |-
    ["ab", "c"]
    ["a", "bc"]
    ["a", "cb"]
    ["ab", "c"]
    ["abc", "d", "defg"]
    ["abcddefg"]
*/

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func Test_1781(t *testing.T) {
	type Data struct {
		word1, word2 []string
	}
	NewTestcases(t).
		Add(true, Data{
			word1: MakeStringSlice(`["ab","c"]`),
			word2: MakeStringSlice(`["a","bc"]`),
		}).
		Add(false, Data{
			word1: MakeStringSlice(`["a","cb"]`),
			word2: MakeStringSlice(`["ab","c"]`),
		}).
		Add(true, Data{
			word1: MakeStringSlice(`["abc","d","defg"]`),
			word2: MakeStringSlice(`["abcddefg"]`),
		}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := arrayStringsAreEqual(input.word1, input.word2)

			a.Equal(td.Expectation, actual)
		})
}
