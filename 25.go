package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.NewInt(1)
	c := new(big.Int)
	var t uint64
	t = 1
	for {
		if len(a.String()) >= 1000 {
			fmt.Println(t)
			break
		}
		c.Set(b)
		b.Add(b, a)
		a.Set(c)
		t++
	}
}
