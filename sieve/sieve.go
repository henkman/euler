package sieve

import (
	"math"
)

type SieveResults []bool

func (sr SieveResults) IsPrime(x uint64) bool {
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}
	return sr[(x-1)/2]
}

/*
	To save space we only store odd numbers, following is an example mapping
	of save spots to their numbers and vice versa
	i: 0 1 2 3 4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
	x: 1 3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51
	->
	x = i*2 +1
	i = (x-1)/2
*/
func Sieve(n uint64) SieveResults {
	var i uint64
	h := n / 2
	sr := make(SieveResults, h)
	for i = 1; i < h; i++ {
		sr[i] = true
	}
	s := uint64(math.Sqrt(float64(n)))
	for i = 1; i <= s; i++ {
		if !sr[i] {
			continue
		}
		x := i*2 + 1         // the actual number
		p := 2 * i * (i + 1) // ((x*x)-1)/2
		// a nice property of the mapping is that you can add a number
		// to the save spot(i) of the number(x) to get the next number
		// that is divisble by the number(x)
		// for example: i = 1, x = 3 -> i + 3 = 4, (4*2+1) = 9
		// for example: i = 3, x = 7 -> i + 7 = 10, (10*2+1) = 21
		for e := p; e < h; e += x {
			sr[e] = false
		}
	}
	return sr
}
