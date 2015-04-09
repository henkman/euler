package main

import (
	"fmt"
)

func main() {
	var s, a, b uint32
	a, b = 1, 2
	for a < 4000000 {
		if a%2 == 0 {
			s += a
		}
		a, b = b, a+b
	}
	fmt.Println(s)
}
