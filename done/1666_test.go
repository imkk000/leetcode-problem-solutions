package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1666"
  title: make-the-string-great
  lang: golang
  type: large
  inputString: |-
    "leEeetcode"
    "abBAcC"
    "s"
*/

func makeGood(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)-1; i++ {
		// skip to remove them
		// when there are not 'a'-'A' and 'A'-'a'
		if b[i]^b[i+1] != 32 {
			continue
		}
		// remove 2 runes
		b = append(b[:i], b[i+2:]...)
		// reset counter (i)
		i = -1
	}
	return string(b)
}

func Test_1666(t *testing.T) {
	NewTestcases(t).
		Add("leetcode", "leEeetcode").
		Add("leetcode", "lEeeetcode").
		Add("ltcode", "lEEeetcode").
		Add("leeeetcode", "leeeetcode").
		Add("ltcode", "lEeeEtcode").
		Add("", "abBAcC").
		Add("s", "s").
		Each(func(a *assert.Assertions, td TestData) {
			actual := makeGood(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
