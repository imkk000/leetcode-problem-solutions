package leetcode_test

import (
	"strconv"
	"testing"

	. "github.com/nattakit-boonyang/go-testcase-builder"
	"github.com/stretchr/testify/assert"
)

/*
metadata:
  id: "412"
  title: "fizz-buzz"
  lang: "golang"
  type: "large"
  stringInput:
  input:
    - "1\n3\n5\n6\n10\n12\n15\n100\n10000"
*/

func Test_412(t *testing.T) {
	NewTestcases(t).
		Add(MakeStringSlice(`["1"]`), 1).
		Add(MakeStringSlice(`["1","2","Fizz"]`), 3).
		Add(MakeStringSlice(`["1","2","Fizz","4","Buzz"]`), 5).
		Add(MakeStringSlice(`["1","2","Fizz","4","Buzz","Fizz"]`), 6).
		Add(MakeStringSlice(`["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz"]`), 10).
		Add(MakeStringSlice(`["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz"]`), 12).
		Add(MakeStringSlice(`["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]`), 15).
		Add(MakeStringSlice(`["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz","16","17","Fizz","19","Buzz","Fizz","22","23","Fizz","Buzz","26","Fizz","28","29","FizzBuzz","31","32","Fizz","34","Buzz","Fizz","37","38","Fizz","Buzz","41","Fizz","43","44","FizzBuzz","46","47","Fizz","49","Buzz","Fizz","52","53","Fizz","Buzz","56","Fizz","58","59","FizzBuzz","61","62","Fizz","64","Buzz","Fizz","67","68","Fizz","Buzz","71","Fizz","73","74","FizzBuzz","76","77","Fizz","79","Buzz","Fizz","82","83","Fizz","Buzz","86","Fizz","88","89","FizzBuzz","91","92","Fizz","94","Buzz","Fizz","97","98","Fizz","Buzz"]`), 100).
		Each(func(a *assert.Assertions, td TestData) {
			actual := fizzBuzz(td.Input.(int))

			a.Equal(td.Expectation, actual)
		})
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzz(1e04)
	}
}

func BenchmarkFizzBuzzP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzzP(1e04)
	}
}

func fizzBuzz(n int) []string {
	s := make([]string, n)
	var fizz, buzz byte
	for i := 1; i <= n; i++ {
		fizz++
		buzz++
		if fizz == 3 && buzz == 5 {
			s[i-1] = "FizzBuzz"
			fizz, buzz = 0, 0
			continue
		}
		if buzz == 5 {
			s[i-1] = "Buzz"
			buzz = 0
			continue
		}
		if fizz == 3 {
			s[i-1] = "Fizz"
			fizz = 0
			continue
		}
		s[i-1] = strconv.Itoa(i)
	}
	return s
}

func fizzBuzzP(n int) []string {
	m := make([]string, 16)
	for i := 0; i <= 15; i++ {
		if i%15 == 0 {
			m[i] = "FizzBuzz"
			continue
		}
		if i%5 == 0 {
			m[i] = "Buzz"
			continue
		}
		if i%3 == 0 {
			m[i] = "Fizz"
			continue
		}
	}
	s := make([]string, n)
	for i := 1; i <= n; i++ {
		if v := m[i%15]; v != "" {
			s[i-1] = v
			continue
		}
		s[i-1] = strconv.Itoa(i)
	}
	return s
}
