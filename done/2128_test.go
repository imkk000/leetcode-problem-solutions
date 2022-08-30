package leetcode_test

import (
	"bytes"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2128"
  title: reverse-prefix-of-word
  lang: golang
  type: large
  inputString: |-
    "abcdefd"
    "d"
    "xyxzxe"
    "z"
    "abcd"
    "z"
    "abc"
    "c"
    "a"
    "a"
    "ab"
    "a"
    "ab"
    "b"
*/

func reversePrefix(word string, ch byte) string {
	b := []byte(word)
	if index := bytes.IndexByte(b, ch); index >= 0 {
		for i := 0; i <= index/2; i++ {
			b[i], b[index-i] = b[index-i], b[i]
		}
	}
	return string(b)
}

func Test_2128(t *testing.T) {
	type Data struct {
		word string
		ch   byte
	}
	NewTestcases(t).
		Add("dcbaefd", Data{"abcdefd", 'd'}).
		Add("zxyxxe", Data{"xyxzxe", 'z'}).
		Add("abcd", Data{"abcd", 'z'}).
		Add("cba", Data{"abc", 'c'}).
		Add("a", Data{"a", 'a'}).
		Add("ab", Data{"ab", 'a'}).
		Add("ba", Data{"ab", 'b'}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := reversePrefix(input.word, input.ch)

			a.Equal(td.Expectation, actual)
		})
}
