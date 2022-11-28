package leetcode_test

import (
	"math"
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

var peekTable = []string{
	"",
	"_1_2_4_8_",
	"_16_23_46_",
	"_128_256_125_",
	"_0124_0248_0469_1289_",
	"_13468_23678_35566_",
	"_011237_122446_224588_",
	"_0145678_0122579_0134449_0368888_",
	"_11266777_23334455_01466788_",
	"_112234778_234455668_012356789_",
	"",
}

func reorderedPowerOf2(n int) bool {
	s := []byte(strconv.Itoa(n))
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	n = int(math.Log10(float64(n))) + 1
	return strings.Contains(peekTable[n], "_"+string(s)+"_")
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
		Add(true, 4096).
		Add(false, int(1e9)).
		Each(func(a *assert.Assertions, td TestData) {
			actual := reorderedPowerOf2(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}
