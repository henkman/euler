package main

import (
	"fmt"
)

func main() {
	const (
		S = 1
		E = 20
	)
	var i, e, max uint64
	max = S
	for i = S + 1; i <= E; i++ {
		max *= i
	}
outer:
	for i = E; i < max; i++ {
		for e = E; e > S; e-- {
			if i%e != 0 {
				continue outer
			}
		}
		fmt.Println(i)
		break
	}
}
