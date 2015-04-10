package main

import (
	"fmt"
	"math"
)

func sieve(n uint64) []bool {
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

func main() {
	const N = 2000000
	p := sieve(N)
	var i, s uint64
	for i = 0; i < N; i++ {
		if p[i] {
			s += i
		}
	}
	fmt.Println(s)
}
