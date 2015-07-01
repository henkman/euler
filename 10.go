package main

import (
	"fmt"
	"github.com/papplampe/euler/sieve"
)

func main() {
	const N = 2000000
	sr := sieve.Sieve(N)
	var i, s uint64
	for i = 0; i < N; i++ {
		if sr.IsPrime(i) {
			s += i
		}
	}
	fmt.Println(s)
}
