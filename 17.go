package main

import (
	"fmt"
)

var (
	digitToWord = []string{
		"",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}
	tennersToWord = []string{
		"",
		"ten",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}
)

func numberToWords(x int) string {
	s := ""
	if x >= 1000 {
		n := x / 1000
		s += digitToWord[n] + "thousand"
	}
	if x >= 100 {
		n := (x % 1000) / 100
		if n != 0 {
			s += digitToWord[n] + "hundred"
		}
	}
	n := x % 100
	if n != 0 {
		if x >= 100 {
			s += "and"
		}
		if n < 20 {
			s += digitToWord[n]
		} else {
			t := n / 10
			if t != 0 {
				s += tennersToWord[t]
			}
			if n%10 != 0 {
				s += digitToWord[n%10]
			}

		}
	}
	return s
}

func main() {
	s := 0
	for i := 1; i <= 1000; i++ {
		s += len(numberToWords(i))
	}
	fmt.Println(s)
}
