package main

import (
	"log"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	log.Println(strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	log.Println(ints)

	s := sort.IntsAreSorted(ints)
	log.Println("Sorted: ", s)
}
