package sieve

import (
	"math"
)

func Sieve(n uint64) []bool {
	var i uint64
	s := uint64(math.Sqrt(float64(n)))
	p := make([]bool, n)
	for i = 2; i < n; i++ {
		p[i] = true
	}
	for i = 2; i <= s; i++ {
		if !p[i] {
			continue
		}
		for e := i * i; e < n; e += i {
			p[e] = false
		}
	}
	return p
}
