package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readTriangle() ([][]uint64, error) {
	file, err := ioutil.ReadFile("p067_triangle.txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(file), "\n")
	triangle := make([][]uint64, len(lines))
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		nums := strings.Split(line, " ")
		triangle[y] = make([]uint64, len(nums))
		for x, sn := range nums {
			n, err := strconv.Atoi(sn)
			if err != nil {
				return nil, err
			}
			triangle[y][x] = uint64(n)
		}
	}

	return triangle, nil
}

func main() {
	triangle, err := readTriangle()
	if err != nil {
		log.Fatal(err)
	}
	for y := 1; y < len(triangle); y++ {
		for x := 0; x < len(triangle[y]); x++ {
			var u uint64
			if x > 0 {
				u = triangle[y-1][x-1]
			}
			if x < len(triangle[y])-1 && triangle[y-1][x] > u {
				u = triangle[y-1][x]
			}
			triangle[y][x] += u
		}
	}

	var l uint64
	for _, n := range triangle[len(triangle)-1] {
		if n > l {
			l = n
		}
	}
	fmt.Println(l)
}
