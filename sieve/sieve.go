package sieve

import (
	"math"
)

func Sieve(n uint64) []bool {
	var i uint64
	s := uint64(math.Sqrt(float64(n)))
	p := make([]bool, n)
	p[2] = true
	for i = 3; i < n; i += 2 {
		p[i] = true
	}
	for i = 3; i <= s; i += 2 {
		if !p[i] {
			continue
		}
		for e := i * i; e < n; e += i {
			p[e] = false
		}
	}
	return p
}
