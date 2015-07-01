package main

import (
	"fmt"
	"github.com/papplampe/euler/sieve"
	"math"
)

func main() {
	x := uint64(600851475143)
	s := uint64(math.Sqrt(float64(x)))
	sr := sieve.Sieve(s)
	for i := s; i > 1; i-- {
		if x%i == 0 && sr.IsPrime(i) {
			fmt.Println(i)
			break
		}
	}
}
