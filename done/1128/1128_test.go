package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1128"
  title: remove-all-adjacent-duplicates-in-string
  lang: golang
  type: large
  inputString: |-
    "abbaca"
    "azxxzy"
    "zxzx"
    "zyyz"
    "z"
    "zy"
    "zz"
*/

func removeDuplicates(s string) string {
	l := len(s)
	st := newStack(l)
	st.push(s[0])
	for i := 1; i < l; i++ {
		if s[i] != st.top() {
			st.push(s[i])
			continue
		}
		st.pop()
	}
	return string(st.data[:st.l])
}

type stack struct {
	data []byte
	l    int
}

func newStack(s int) stack {
	return stack{
		data: make([]byte, s),
		l:    0,
	}
}

func (s *stack) push(r byte) {
	s.data[s.l] = r
	s.l++
}

func (s *stack) pop() {
	if s.empty() {
		return
	}
	s.l--
}

func (s *stack) empty() bool {
	return s.l == 0
}

func (s *stack) top() byte {
	if s.empty() {
		return 0
	}
	return s.data[s.l-1]
}

func Test_1128(t *testing.T) {
	NewTestcases(t).
		Add("ca", "abbaca").
		Add("ay", "azxxzy").
		Add("zxzx", "zxzx").
		Add("", "zyyz").
		Add("z", "z").
		Add("zy", "zy").
		Each(func(a *assert.Assertions, td TestData) {
			actual := removeDuplicates(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
