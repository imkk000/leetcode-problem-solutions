package leetcode_test

import (
	"testing"
	"time"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1289"
  title: day-of-the-week
  lang: golang
  type: large
  inputString: |-
    31
    8
    2019
    18
    7
    1999
    15
    8
    1993
    31
    12
    2100
    1
    1
    1971
*/

func dayOfTheWeek(day, month, year int) string {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return t.Format("Monday")
}

func Test_1289(t *testing.T) {
	newTestcase := func(d, m, y int) TestDataFunc {
		newExpected := func(d, m, y int) string {
			t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
			return t.Format("Monday")
		}
		return func() TestData {
			return TestData{
				Expectation: newExpected(d, m, y),
				Input:       []int{d, m, y},
			}
		}
	}

	NewTestcases(t).
		AddFunc(newTestcase(1, 1, 1971)).
		AddFunc(newTestcase(2, 1, 1971)).
		AddFunc(newTestcase(31, 1, 1971)).
		AddFunc(newTestcase(1, 2, 1971)).
		AddFunc(newTestcase(28, 2, 1971)).
		AddFunc(newTestcase(1, 3, 1971)).
		AddFunc(newTestcase(31, 12, 1971)).
		AddFunc(newTestcase(1, 1, 1972)).
		AddFunc(newTestcase(2, 1, 1972)).
		AddFunc(newTestcase(31, 1, 1972)).
		AddFunc(newTestcase(1, 2, 1972)).
		AddFunc(newTestcase(28, 2, 1972)).
		AddFunc(newTestcase(29, 2, 1972)).
		AddFunc(newTestcase(1, 3, 1972)).
		AddFunc(newTestcase(31, 12, 1972)).
		AddFunc(newTestcase(1, 1, 1973)).
		AddFunc(newTestcase(31, 8, 2019)).
		AddFunc(newTestcase(18, 7, 1999)).
		AddFunc(newTestcase(15, 8, 1993)).
		AddFunc(newTestcase(31, 12, 2100)).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.([]int)
			day, month, year := input[0], input[1], input[2]
			actual := dayOfTheWeek(day, month, year)

			a.Equal(td.Expectation, actual)
		})
}
