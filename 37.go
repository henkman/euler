package main

import (
	"fmt"
	"github.com/papplampe/euler/primegen"
	"strconv"
)

var (
	pg *primegen.Generator
)

func isTruncatablePrime(p uint64) bool {
	s := []byte(fmt.Sprint(p))
	n := len(s)
	for i := 1; i < n; i++ {
		l, _ := strconv.ParseUint(string(s[i:]), 10, 64)
		if !pg.IsPrime(l) {
			return false
		}
		r, _ := strconv.ParseUint(string(s[:i]), 10, 64)
		if !pg.IsPrime(r) {
			return false
		}
	}
	return true
}

func main() {
	pg = primegen.New(50000)
	var sum uint64
	// 23, 37, 53, 73 already in primegen
	sum = 23 + 37 + 53 + 73
	n := 4
	for n < 11 {
		p := pg.NextAdd()
		if isTruncatablePrime(p) {
			sum += p
			n++
		}
	}
	fmt.Println(sum)
}
