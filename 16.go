package main

/*
215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 21000?
*/

import (
	"fmt"
	"math/big"
)

func main() {
	n := new(big.Int)
	n.SetUint64(2)
	n.Exp(n, big.NewInt(1000), nil)
	var s uint64
	for _, c := range n.String() {
		s += uint64(c - '0')
	}
	fmt.Println(s)
}
