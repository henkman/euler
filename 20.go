package main

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
