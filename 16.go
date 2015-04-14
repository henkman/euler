package main

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
