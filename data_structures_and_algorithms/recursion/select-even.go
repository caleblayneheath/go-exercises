package main

import (
	"fmt"
)

func selectEven(slice []int) []int {
	result := []int{}

	if len(slice) == 0 {
		return result
	}

	if slice[0]%2 == 0 {
		result = append(result, slice[0])
	}

	if len(slice) > 1 {
		result = append(result, selectEven(slice[1:])...)
	}

	return result
}

func main() {
	slc1 := []int{0, 1, 2, 3, 4, 5}
	slc2 := []int{}

	fmt.Println(selectEven(slc1)) // 0 2 4
	fmt.Println(selectEven(slc2)) // 0

}

