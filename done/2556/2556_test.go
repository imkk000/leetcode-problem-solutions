package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2556"
  title: convert-the-temperature
  lang: golang
  type: large
  inputString: |-
    36.50
    122.11
    0
    1000
*/

func convertTemperature(c float64) []float64 {
	return []float64{c + 273.15, c*1.80 + 32}
}

func Test_2556(t *testing.T) {
	NewTestcases(t).
		Add([]float64{309.65000, 97.70000}, float64(36.50)).
		Add([]float64{395.26000, 251.79800}, float64(122.11)).
		Add([]float64{273.15000, 32.00000}, float64(0)).
		Add([]float64{1273.15000, 1832.00000}, float64(1000)).
		Each(func(a *assert.Assertions, td TestData) {
			actual := convertTemperature(td.Input.(float64))

			a.Equal(td.Expectation, actual)
		})
}
