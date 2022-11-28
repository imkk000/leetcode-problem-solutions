package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "393"
  title: utf-8-validation
  lang: golang
  type: large
  inputString: |-
    [197,130,1]
    [235,140,4]
    [194,155,231,184,185,246,176,131,161,222,174,227,162,134,241,154,168,185,218,178,229,187,139,246,178,187,139,204,146,225,148,179,245,139,172,134,193,156,233,131,154,240,166,188,190,216,150,230,145,144,240,167,140,163,221,190,238,168,139,241,154,159,164,199,170,224,173,140,244,182,143,134,206,181,227,172,141,241,146,159,170,202,134,230,142,163,244,172,140,191]
    [145]
    [248,130,130,130]
*/

func validUtf8(nums []int) bool {
	validNextBits := func(v int) bool {
		return v >= 128 && v <= 191
	}
	var next int
	for _, v := range nums {
		if next == 0 {
			if validNextBits(v) || v >= 248 {
				return false
			}
			if v <= 127 {
				next = 0
			} else if v <= 223 {
				next = 1
			} else if v <= 239 {
				next = 2
			} else if v <= 247 {
				next = 3
			}
			continue
		}
		next--
		if !validNextBits(v) {
			return false
		}
	}
	return next == 0
}

func Test_393(t *testing.T) {
	NewTestcases(t).
		Add(true, MakeIntSlice("[197,130,1]")).
		Add(false, MakeIntSlice("[235,140,4]")).
		Add(true, MakeIntSlice("[194,155,231,184,185,246,176,131,161,222,174,227,162,134,241,154,168,185,218,178,229,187,139,246,178,187,139,204,146,225,148,179,245,139,172,134,193,156,233,131,154,240,166,188,190,216,150,230,145,144,240,167,140,163,221,190,238,168,139,241,154,159,164,199,170,224,173,140,244,182,143,134,206,181,227,172,141,241,146,159,170,202,134,230,142,163,244,172,140,191]")).
		Add(false, MakeIntSlice("[145]")).
		Add(false, MakeIntSlice("[248,130,130,130]")).
		Each(func(a *assert.Assertions, td TestData) {
			actual := validUtf8(td.Input.([]int))

			a.Equal(td.Expectation, actual)
		})
}
