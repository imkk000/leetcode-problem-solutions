package leetcode_test

import (
	"regexp"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "520"
  title: detect-capital
  lang: golang
  type: large
  inputString: |-
    "USA"
    "FlaG"
    "leetcode"
    "Leetcode"
    "a"
    "A"
    "AnT"
*/

func detectCapitalUse(word string) bool {
	return regexp.
		MustCompile("^(([A-Z]+)|([A-Z]?[a-z]*))$").
		MatchString(word)
}

func Test_520(t *testing.T) {
	NewTestcases(t).
		Add(true, "USA").
		Add(true, "leetcode").
		Add(true, "Leetcode").
		Add(true, "a").
		Add(true, "A").
		Add(false, "AnT").
		Each(func(a *assert.Assertions, td TestData) {
			actual := detectCapitalUse(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
