package leetcode_test

import (
	"strings"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "38"
  title: count-and-say
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
    8
    9
    10
    11
    12
    13
    14
    15
    16
    17
    18
    19
    20
    21
    22
    23
    24
    25
    26
    27
    28
    29
    30
*/

func countAndSay(n int) (s string) {
	s = "1"
	if n <= 1 {
		return
	}

	var v strings.Builder
	start, end, l := 0, 1, 1
	for i := 2; i <= n; i++ {
		for end < l {
			if s[start] != s[end] {
				v.WriteByte(byte((end - start) + '0'))
				v.WriteByte(s[start])
				start = end
			}
			end++
		}
		if start < end {
			v.WriteByte(byte((end - start) + '0'))
			v.WriteByte(s[start])
		}
		s = v.String()
		l = v.Len()
		v.Reset()
		start = 0
		end = 1
	}
	return
}

func Test_38(t *testing.T) {
	NewTestcases(t).
		Add("1", 1).
		Add("11", 2).
		Add("21", 3).
		Add("1211", 4).
		Each(func(a *assert.Assertions, td TestData) {
			actual := countAndSay(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
