package main

import (
	"fmt"
)

func isDivisableByRange(x, s, e uint64) bool {
	for i := e; i >= s; i-- {
		if x%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	const (
		S = 1
		E = 20
	)
	var i, max uint64
	max = S
	for i = S + 1; i <= E; i++ {
		max *= i
	}
	for i = E; i < max; i++ {
		if isDivisableByRange(i, S+1, E) {
			fmt.Println(i)
			break
		}
	}
}
