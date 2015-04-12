package main

import (
	"fmt"
)

func main() {
	const (
		W = 20
		H = 20
	)
	var grid [H + 1][W + 1]uint64

	for y := 0; y < H; y++ {
		grid[y][W] = 1
	}
	for x := 0; x < W; x++ {
		grid[H][x] = 1
	}

	for y := H - 1; y >= 0; y-- {
		for x := W - 1; x >= 0; x-- {
			grid[y][x] = grid[y+1][x] + grid[y][x+1]
		}
	}
	fmt.Println(grid[0][0])
}
