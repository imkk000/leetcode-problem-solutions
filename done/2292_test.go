package leetcode_test

import (
	_ "embed"
	"strings"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "2292"
  title: counting-words-with-a-given-prefix
  lang: golang
  type: large
  inputString: |-
    ["pay","attention","practice","attend"]
    "at"
    ["leetcode","win","loops","success"]
    "code"
*/

func prefixCount(words []string, pref string) (c int) {
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			c++
		}
	}
	return
}

func Test_2292(t *testing.T) {
	type Data struct {
		words  []string
		prefix string
	}
	NewTestcases(t).
		Add(2, Data{MakeStringSlice(`["pay","attention","practice","attend"]`), "at"}).
		Add(0, Data{MakeStringSlice(`["leetcode","win","loops","success"]`), "code"}).
		Add(1, Data{MakeStringSlice(input), "cogqrmskvhpmheiacxzknhrlrbpqtbhkvbeadypxhupihjfpncsghfufkpsvqwyefazucdzowasivwuytlpvqfvxaxwxvslmtxqj"}).
		Add(0, Data{[]string{"a"}, "abc"}).
		Each(func(a *assert.Assertions, td TestData) {
			input := td.Input.(Data)
			actual := prefixCount(input.words, input.prefix)

			a.Equal(td.Expectation, actual)
		})
}

func BenchmarkIterative(b *testing.B) {
	in := MakeStringSlice(input)
	for i := 0; i < b.N; i++ {
		prefixCount(in, "at")
	}
}

// go:exclude
//
//go:embed testdata/2292_1_in
var input string
