package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1448"
  title: maximum-69-number
  lang: golang
  type: large
  inputString: |-
    9669
    9996
    9999
    9
    69
    66
    6666
*/

func maximum69Number(num int) int {
	if 6 == num/1000 {
		return num + 3000
	}
	if 6 == (num%1000)/100 {
		return num + 300
	}
	if 6 == ((num%1000)%100)/10 {
		return num + 30
	}
	if 6 == num%10 {
		return num + 3
	}
	return num
}

func Test_1448(t *testing.T) {
	NewTestcases(t).
		Add(9969, 9669).
		Add(9999, 9996).
		Add(9999, 9999).
		Add(9, 6).
		Add(99, 69).
		Add(96, 66).
		Add(9666, 6666).
		Each(func(a *assert.Assertions, td TestData) {
			actual := maximum69Number(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
