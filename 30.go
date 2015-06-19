package main

import (
	"fmt"
	"math"
)

func sumOfPowerOfDigits(n, power uint64) uint64 {
	var s uint64
	t := fmt.Sprint(n)
	for _, c := range []byte(t) {
		s += uint64(math.Pow(float64(c-'0'), float64(power)))
	}
	return s
}

func main() {
	var i, sp, s uint64
	for i = 2; i < 200000; i++ {
		sp = sumOfPowerOfDigits(i, 5)
		if i == sp {
			s += i
		}
	}
	fmt.Println(s)
}
