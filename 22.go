package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

/*
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names,
begin by sorting it into alphabetical order.
Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list.
So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

func main() {
	file, err := ioutil.ReadFile("./p022_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	names := strings.Split(string(file), ",")
	for i, name := range names {
		names[i] = name[1 : len(name)-1]
	}
	sort.Strings(names)
	var s uint64
	for i, name := range names {
		ns := 0
		for _, c := range name {
			ns += int(c - 'A' + 1)
		}
		s += uint64(ns * (i + 1))
	}
	fmt.Println(s)
}
