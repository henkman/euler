package main

import (
	"euler/sieve"
	"fmt"
)

func main() {
	const N = 2000000
	p := sieve.Sieve(N)
	var i, s uint64
	for i = 0; i < N; i++ {
		if p[i] {
			s += i
		}
	}
	fmt.Println(s)
}
