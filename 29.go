package main

import (
	"fmt"
	"math/big"
)

func main() {
	const (
		S = 2
		E = 100
	)
	n := make(map[string]interface{})
	var a, b uint64
	ab := new(big.Int)
	bb := new(big.Int)
	t := new(big.Int)
	for a = S; a <= E; a++ {
		for b = S; b <= E; b++ {
			ab.SetUint64(a)
			bb.SetUint64(b)
			t.Exp(ab, bb, nil)
			n[t.String()] = true
		}
	}
	fmt.Println(len(n))
}
