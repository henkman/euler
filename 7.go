package main

import (
	"fmt"
	"github.com/papplampe/euler/primegen"
)

func main() {
	const N = 10001
	pg := primegen.New(N)
	for pg.Len() < N {
		pg.NextAdd()
	}
	fmt.Println(pg.Last())
}
