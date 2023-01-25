package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1920"
  title: determine-color-of-a-chessboard-square
  lang: golang
  type: large
  inputString: |-
    "a1"
    "h3"
    "c7"
*/

func squareIsWhite(c string) bool {
	return (((c[0] - 'a' + 1) & 1) ^ ((c[1] - '0') & 1)) == 1
}

func Test_1920(t *testing.T) {
	NewTestcases(t).
		Add(false, "a1").
		Add(true, "h3").
		Add(false, "c7").
		Add(false, "h8").
		Add(true, "a8").
		Add(true, "h1").
		Each(func(a *assert.Assertions, td TestData) {
			actual := squareIsWhite(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
