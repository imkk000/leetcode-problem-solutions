package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "742"
  title: to-lower-case
  lang: golang
  type: large
  inputString: |-
    "Hello"
    "here"
    "LOVELY"
    "A"
    "a"
*/

func toLowerCase(s string) string {
	r := []byte(s)
	for i, c := range r {
		if c >= 'A' && c <= 'Z' {
			r[i] += 32
		}
	}
	return string(r)
}

func Test_742(t *testing.T) {
	NewTestcases(t).
		Add("hello", "Hello").
		Add("here", "here").
		Add("lovely", "LOVELY").
		Add("a", "A").
		Add("a", "a").
		Add("al&phabet", "al&phaBET").
		Each(func(a *assert.Assertions, td TestData) {
			actual := toLowerCase(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
