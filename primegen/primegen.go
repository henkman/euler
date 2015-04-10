package primegen

import (
	"fmt"
	"io"
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
outer:
	for i := lp + 2; ; i++ {
		for e := 0; e < p.Len(); e++ {
			if i%p.Primes[e] == 0 {
				continue outer
			}
		}
		return i
	}
	panic("whoops")
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
	for {
		x := p.Next()
		if x > n {
			break
		}
		p.Add(x)
	}
}

func (p *Generator) GenerateNext() {
	p.Add(p.Next())
}

func (p *Generator) Print(out io.Writer) {
	for i := 0; i < p.Len(); i++ {
		fmt.Fprintln(out, p.Primes[i])
	}
}
