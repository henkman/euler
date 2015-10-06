package main

import (
	"fmt"
	"math"
)

const (
	MAX = 28123
)

func sumOfDivisors(x uint64) uint64 {
	var i, s, o, sum uint64
	s = uint64(math.Sqrt(float64(x)))
	for i = 2; i <= s; i++ {
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
	var i uint64
	abundent := make([]uint64, 0, 1000)
	for i = 2; i < MAX; i++ {
		d := sumOfDivisors(i)
		if d > i {
			abundent = append(abundent, i)
		}
	}

	sumOfAbundand := make([]bool, MAX+1)
	for ia, a := range abundent {
		for ib := ia; ib < len(abundent); ib++ {
			b := abundent[ib]
			sum := a + b
			if sum <= MAX {
				sumOfAbundand[sum] = true
			} else {
				break
			}
		}
	}

	var s uint64
	for i = 1; i <= MAX; i++ {
		if !sumOfAbundand[i] {
			s += i
		}
	}
	fmt.Println(s)
}
