package leetcode_test

import (
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "1472"
  title: increasing-decreasing-string
  lang: golang
  type: large
  inputString: |-
    "aaaabbbbcccc"
    "abccbaabccba"
    "rat"
    "mvydjbsvivswyievhuscgtvtzrwqfxywglckariatotljnucpotyczjbcaqxiuhiwanvsigjucbzkjwxsjprhljsjqjqcfazdtusxkxfkdxsrombwfdfpcjqcxeqgnlajblvkyjluizbkboyfmxspnvthcpxircscuqbhdpdiliqbluhmrplyqkbqrvjypuydumlkhjflesbnjsejtwafvuhogtqynpnbvmoqwjahmpaybmcrxmnotvmpjqdcjwcsnspipuidnotqtijwuyoidhrvyjgliczgourirpaexiwmlbvamyqtghiflglzoczinahusnwnmjujkortyfjcegwtuvieoreczgrmqlwblpxflwrnoizdskvdgbcrgpjzfqzilyingxwxiqsjqmfahezwxnwhdzmqmsltjdkxenresdakrkqibgmgyviebjnueabcibglzdqlgrcfvuxvavornxqkabqpbsartflwbuzckuhenfo"
*/

func sortString(s string) string {
	b := []byte(s)
	m := make([]int, 26)
	for _, r := range b {
		m[r-'a']++
	}
	l := len(b)
	for c := 0; c < l; {
		for j := 0; j < 26; j++ {
			if m[j] > 0 {
				b[c] = byte(j + 'a')
				m[j]--
				c++
			}
		}
		for j := 25; j >= 0; j-- {
			if m[j] > 0 {
				b[c] = byte(j + 'a')
				m[j]--
				c++
			}
		}
	}
	return string(b)
}

func Test_1472(t *testing.T) {
	NewTestcases(t).
		Add("abccbaabccba", "aaaabbbbcccc").
		Add("abccbaabccba", "abccbaabccba").
		Add("art", "rat").
		Add("abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcbaabcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcbaabcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcbaabcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcbaabcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcbaabcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcbaabcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcbaabcdfghijklmnopqrstuvwxyzzyxwvutsrqponmljigfdcbaabcgijlmnpqrstuvwxyxwvusrqnmljigcbabcijlnqrsuvwvusrqnljicbbcijlqruqljicbbcijlqqjiijjiijj", "mvydjbsvivswyievhuscgtvtzrwqfxywglckariatotljnucpotyczjbcaqxiuhiwanvsigjucbzkjwxsjprhljsjqjqcfazdtusxkxfkdxsrombwfdfpcjqcxeqgnlajblvkyjluizbkboyfmxspnvthcpxircscuqbhdpdiliqbluhmrplyqkbqrvjypuydumlkhjflesbnjsejtwafvuhogtqynpnbvmoqwjahmpaybmcrxmnotvmpjqdcjwcsnspipuidnotqtijwuyoidhrvyjgliczgourirpaexiwmlbvamyqtghiflglzoczinahusnwnmjujkortyfjcegwtuvieoreczgrmqlwblpxflwrnoizdskvdgbcrgpjzfqzilyingxwxiqsjqmfahezwxnwhdzmqmsltjdkxenresdakrkqibgmgyviebjnueabcibglzdqlgrcfvuxvavornxqkabqpbsartflwbuzckuhenfo").
		Add("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa").
		Add("abbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabba", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb").
		Each(func(a *assert.Assertions, td TestData) {
			actualList := []string{
				sortString(td.Input.(string)),
			}

			for _, actual := range actualList {
				a.Equal(td.Expectation, actual)
			}
		})
}
