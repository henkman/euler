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
	x := uint64(600851475143)
	s := uint64(math.Sqrt(float64(x)))
	p := sieve(s)
	for i := s; i > 1; i-- {
		if x%i == 0 && p[i] {
			fmt.Println(i)
			break
		}
	}
}
