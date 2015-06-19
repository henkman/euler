package main

import (
	"fmt"
	"strconv"
)

const (
	N = 1000000
)

func isPalindrome(s []byte) bool {
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

func toBinary(n uint64) []byte {
	return []byte(strconv.FormatUint(n, 2))
}

func main() {
	var i, s uint64
	for i = 1; i < N; i++ {
		if isPalindrome([]byte(fmt.Sprint(i))) && isPalindrome(toBinary(i)) {
			s += i
		}
	}
	fmt.Println(s)
}
