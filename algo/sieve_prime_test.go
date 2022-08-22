package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func LoopPrime(n int) []int {
	isPrime := func(n int) bool {
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	var s []int
	for i := 2; i <= n; i++ {
		if !isPrime(i) {
			continue
		}
		s = append(s, i)
	}
	return s
}

func TestLoopPrime(t *testing.T) {
	expectation := SievePrime(1e4)

	actual := LoopPrime(1e4)

	assert.Equal(t, expectation, actual)
}

func BenchmarkLoopPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopPrime(1e4)
	}
}

func BenchmarkSievePrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SievePrime(1e4)
	}
}

func SievePrime(n int) []int {
	tb := make([]bool, n+1)
	tb[1] = true
	for i := 2; i <= n; i++ {
		if tb[i] {
			continue
		}
		for j := i << 1; j <= n; j += i {
			tb[j] = true
		}
	}
	var s []int
	for i := 1; i <= n; i++ {
		if tb[i] {
			continue
		}
		s = append(s, i)
	}
	return s
}
