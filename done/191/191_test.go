package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "191"
  title: number-of-1-bits
  lang: golang
  type: large
  inputString: |-
    00000000000000000000000000001011
    00000000000000000000000010000000
    11111111111111111111111111111101
*/

var pTable = []byte{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}

func hammingWeight(n uint32) int {
	return int(pTable[n>>28] + pTable[(n>>24)&0xf] +
		pTable[n>>20&0xf] + pTable[n>>16&0xf] +
		pTable[n>>12&0xf] + pTable[n>>8&0xf] +
		pTable[n>>4&0xf] + pTable[n&0xf])
}

func Test_191(t *testing.T) {
	NewTestcases(t).
		Add(3, uint32(0b00000000000000000000000000001011)).
		Add(1, uint32(0b00000000000000000000000010000000)).
		Add(31, uint32(0b11111111111111111111111111111101)).
		Each(func(a *assert.Assertions, td TestData) {
			actual := hammingWeight(td.Input.(uint32))

			a.Equal(td.Expectation, actual)
		})
}
