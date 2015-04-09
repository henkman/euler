package main

import (
	"fmt"
)

func main() {
	var susq, sqsu, i uint32
	for i = 1; i <= 100; i++ {
		susq += i * i
		sqsu += i
	}
	sqsu *= sqsu
	fmt.Println(sqsu - susq)
}
