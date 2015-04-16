package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

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
