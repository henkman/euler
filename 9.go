package main

import (
	"fmt"
)

func main() {
	const N = 1000
	for a := 1; a < N/3; a++ {
		for b := a + 1; b < N/2; b++ {
			c := N - b - a
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
