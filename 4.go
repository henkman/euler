package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)
	h := len(s) / 2
	if h%2 != 0 {
		h += 1
	}
	for i, e := 0, len(s)-1; i < h; i, e = i+1, e-1 {
		if s[i] != s[e] {
			return false
		}
	}
	return true
}

func main() {
	l := 0
	for i := 999; i > 0; i-- {
		for e := i - 1; e > 0; e-- {
			p := i * e
			if isPalindrome(p) && p > l {
				l = p
			}
		}
	}
	fmt.Println(l)
}
