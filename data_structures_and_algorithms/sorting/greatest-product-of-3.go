package main

import (
	"fmt"
	"sort"
)

// assumes non-empty slice
func greatestProductOf3(slc []int) int {
	slc = append([]int{}, slc...)
	sort.Ints(slc)
	
	var product int
	
	switch len(slc) {
	case 1:
		product = slc[0]
	case 2:
		product = slc[0] * slc[1]
	default:
		slc = slc[len(slc)-3:]
		product = slc[0] * slc[1] * slc[2]
	}

	return product
}

func main() {
	slc1 := []int{4, 2, 1, 3}
	fmt.Println(greatestProductOf3(slc1))
}
