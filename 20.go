package main

/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

import (
	"fmt"
	"math/big"
)

func main() {
	fac := big.NewInt(1)
	for i := 2; i < 100; i++ {
		mul := big.NewInt(int64(i))
		fac.Mul(fac, mul)
	}
	var sum uint64
	for _, c := range fac.String() {
		sum += uint64(c - '0')
	}
	fmt.Println(sum)
}
