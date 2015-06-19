package main

import (
	"fmt"
	"github.com/papplampe/euler/sieve"
	"strconv"
)

var (
	primes []bool
)

const (
	N = 1000000
)

func rotateBytes(s []byte) {
	last := s[len(s)-1]
	for i := len(s) - 1; i > 0; i-- {
		s[i] = s[i-1]
	}
	s[0] = last
}

func isCircularPrime(n uint64) bool {
	if !primes[n] {
		return false
	}
	t := []byte(fmt.Sprint(n))
	if len(t) == 1 {
		return true
	}
	for i := 0; i < len(t); i++ {
		rotateBytes(t)
		x, _ := strconv.Atoi(string(t))
		if !primes[x] {
			return false
		}
	}
	return true
}

func main() {
	var i, s uint64
	primes = sieve.Sieve(N)
	for i = 1; i < N; i++ {
		if isCircularPrime(i) {
			s++
		}
	}
	fmt.Println(s)
}
