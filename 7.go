package main

import (
	"fmt"
)

type Primes []uint64

func (p Primes) NextPrime() uint64 {
	lp := p[len(p)-1]
outer:
	for i := lp + 2; ; i++ {
		for e := 0; e < len(p); e++ {
			if i%p[e] == 0 {
				continue outer
			}
		}
		return i
	}
	panic("whoops")
}

func main() {
	primes := make(Primes, 0, 10001)
	primes = append(primes, []uint64{
		2, 3, 5, 7, 11, 13,
	}...)
	for len(primes) < 10001 {
		np := primes.NextPrime()
		primes = append(primes, np)
	}
	fmt.Println(primes[len(primes)-1])
}
