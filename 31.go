package main

import (
	"fmt"
)

func numberOfCoinCombinations(n uint, coins []uint) uint {
	m := make([]uint, n+1)
	m[0] = 1
	for _, coin := range coins {
		for t := coin; t <= n; t++ {
			m[t] += m[t-coin]
		}
	}
	return m[n]
}

func main() {
	coins := []uint{
		200, 100, 50, 20, 10, 5, 2, 1,
	}
	fmt.Println(numberOfCoinCombinations(200, coins))
}
