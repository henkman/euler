package primegen

import (
	"fmt"
	"io"
	"math"
)

type Generator struct {
	Primes []uint64
}

func New(capacity uint64) *Generator {
	p := new(Generator)
	p.Primes = make([]uint64, 0, capacity)
	p.Primes = append(p.Primes, []uint64{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41,
		43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
	}...)
	return p
}

func (p *Generator) Next() uint64 {
	lp := p.Last()
	m := uint64(math.Sqrt(float64(lp)))
	egp := p.Primes
	for i := 0; i < p.Len(); i++ {
		if p.Primes[i] > m {
			egp = p.Primes[:i]
			break
		}
	}
outer:
	for i := lp + 2; ; i += 2 {
		for _, p := range egp {
			if i%p == 0 {
				continue outer
			}
		}
		return i
	}
}

func (p *Generator) IsPrime(x uint64) bool {
	if (x != 2 && x%2 == 0) || (x != 3 && x%3 == 0) {
		return false
	}
	p.GenerateUpTo(x)
	for _, p := range p.Primes {
		if p == x {
			return true
		}
	}
	return false
}

func (p *Generator) Add(x uint64) {
	p.Primes = append(p.Primes, x)
}

func (p *Generator) Last() uint64 {
	return p.Primes[len(p.Primes)-1]
}

func (p *Generator) Len() int {
	return len(p.Primes)
}

func (p *Generator) GenerateUpTo(n uint64) {
	x := p.Next()
	for x < n {
		p.Add(x)
		x = p.Next()
	}
}

func (p *Generator) NextAdd() uint64 {
	n := p.Next()
	p.Add(n)
	return n
}

func (p *Generator) Print(out io.Writer) {
	for i := 0; i < p.Len(); i++ {
		fmt.Fprintln(out, p.Primes[i])
	}
}
