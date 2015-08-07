package main

import (
	"fmt"
	"github.com/papplampe/euler/primegen"
	"math"
)

var (
	pg *primegen.Generator
)

func divisors(x uint64) uint64 {
	var e uint64
	r := x
	d := uint64(1)
	s := uint64(math.Sqrt(float64(x)))
	pg.GenerateUpTo(s)
	for _, p := range pg.Primes {
		if p*p > x {
			return d * 2
		}
		e = 1
		for r%p == 0 {
			e++
			r /= p
		}
		d *= e
		if r == 1 {
			break
		}
	}
	return d
}

func main() {
	var i, t, d uint64
	pg = primegen.New(100)
	for i, t = 1, 1; d < 500; i, t = i+1, t+i+1 {
		d = divisors(t)
	}
	fmt.Println(t - i)
}
