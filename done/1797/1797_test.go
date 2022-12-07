package leetcode_test

import (
	"strings"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1797"
  title: goal-parser-interpretation
  lang: golang
  type: large
  inputString: |-
    "G()(al)"
    "G()()()()(al)"
    "(al)G(al)()()G"
*/

func interpret(command string) string {
	return strings.
		NewReplacer("()", "o", "(al)", "al").
		Replace(command)
}

func Test_1797(t *testing.T) {
	NewTestcases(t).
		Add("Goal", "G()(al)").
		Add("Gooooal", "G()()()()(al)").
		Add("alGalooG", "(al)G(al)()()G").
		Each(func(a *assert.Assertions, td TestData) {
			actual := interpret(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
