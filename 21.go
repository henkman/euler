package main

import (
	"fmt"
	"math"
)

func sumOfDivisors(x uint64) uint64 {
	var i, s, o, sum uint64
	s = uint64(math.Sqrt(float64(x)))
	for i = 2; i < s; i++ {
		if x%i == 0 {
			o = x / i
			if o == i {
				sum += i
			} else {
				sum += i + o
			}
		}
	}
	return 1 + sum
}

func main() {
	const N = 10000
	divs := make(map[uint64]uint64, N-1)
	var i, s uint64
	for i = 2; i < N; i++ {
		divs[i] = sumOfDivisors(i)
	}
	for i, d := range divs {
		if o, ok := divs[d]; ok && i == o && i != d {
			s += i
		}
	}
	fmt.Println(s)
}
