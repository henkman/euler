package main

import (
	"euler/primegen"
	"fmt"
)

func main() {
	const N = 10001
	pg := primegen.New(N)
	for pg.Len() < N {
		pg.GenerateNext()
	}
	fmt.Println(pg.Last())
}
