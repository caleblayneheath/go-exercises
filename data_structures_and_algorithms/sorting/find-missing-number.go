package main

import (
	"fmt"
	"sort"
)

// given an array of ints, 0 to n,
// with one int less than n missing,
// return that missing int

// assumes non-empty slice
func findMissingNumber(slc []int) int {
	slc = append([]int{}, slc...)
	sort.Ints(slc)
	
	missing := -1
	for i, num := range(slc) {
		if (i != num) {
			missing = i
			break
		}
	}	

	return missing
}

func main() {
	slc1 := []int{4, 2, 1, 0}
	fmt.Println(findMissingNumber(slc1))
}

