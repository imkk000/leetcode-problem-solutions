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

func removeDuplicates1128(s string) string {
	l := len(s)
	st := newStack1128(l)
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

type stack1128 struct {
	data []byte
	l    int
}

func newStack1128(s int) stack1128 {
	return stack1128{
		data: make([]byte, s),
		l:    0,
	}
}

func (s *stack1128) push(r byte) {
	s.data[s.l] = r
	s.l++
}

func (s *stack1128) pop() {
	if s.empty() {
		return
	}
	s.l--
}

func (s *stack1128) empty() bool {
	return s.l == 0
}

func (s *stack1128) top() byte {
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
			actual := removeDuplicates1128(td.Input.(string))

			a.Equal(td.Expectation, actual)
		})
}
