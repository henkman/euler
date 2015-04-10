package main

import (
	"euler/sieve"
	"fmt"
	"math"
)

func main() {
	x := uint64(600851475143)
	s := uint64(math.Sqrt(float64(x)))
	p := sieve.Sieve(s)
	for i := s; i > 1; i-- {
		if x%i == 0 && p[i] {
			fmt.Println(i)
			break
		}
	}
}
