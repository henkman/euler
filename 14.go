package main

import (
	"fmt"
)

type Result struct {
	Chain uint64
	Start uint64
}

var (
	ks = map[uint64]uint64{}
)

func main() {
	const N = 1000000
	var r Result
	var i, c, n uint64
	for i = 13; i < N; i++ {
		c = 1
		for n = i; n != 1; {
			if k, ok := ks[n]; ok {
				c += k
				break
			}
			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}
			c++
		}
		ks[i] = c
		if c > r.Chain {
			r.Chain = c
			r.Start = i
		}
	}
	fmt.Println(r.Start)
}
