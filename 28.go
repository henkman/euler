package main

import (
	"fmt"
)

func main() {
	const N = 1001
	n := 1
	l := 2
	s := 1
	for i := 1; n < (N * N); i++ {
		n += l
		if i%4 == 0 {
			l += 2
		}
		s += n
	}
	fmt.Println(s)
}
