package leetcode_test

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "900"
  title: reordered-power-of-2
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
    16
    61
    32
    23
    64
    46
    100
    1000
    10000
    100000
    1000000
    10000000
    100000000
    1000000000
*/

const pTable = ":1:2:4:8:16:23:46:128:256:125:0124:0248:0469:1289:13468:23678:35566:011237:122446:224588:0145678:0122579:0134449:0368888:11266777:23334455:01466788:112234778:234455668:012356789:"

func reorderedPowerOf2(n int) bool {
	s := []byte(strconv.Itoa(n))
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return strings.Contains(pTable, ":"+string(s)+":")
}

func Test_900(t *testing.T) {
	NewTestcases(t).
		Add(true, 1).
		Add(true, 2).
		Add(false, 3).
		Add(true, 4).
		Add(false, 5).
		Add(false, 6).
		Add(false, 7).
		Add(true, 8).
		Add(false, 9).
		Add(false, 10).
		Add(false, 11).
		Add(true, 16).
		Add(true, 61).
		Add(true, 32).
		Add(true, 23).
		Add(true, 64).
		Add(true, 46).
		Each(func(a *assert.Assertions, td TestData) {
			actual := reorderedPowerOf2(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
