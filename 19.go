package main

import (
	"fmt"
	"time"
)

func main() {
	sundays := 0
	for y := 1901; y <= 2000; y++ {
		for m := time.January; m <= time.December; m++ {
			date := time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
			if date.Weekday() == time.Sunday {
				sundays++
			}
		}
	}
	fmt.Println(sundays)
}
