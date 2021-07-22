/*
input: two arrays, each with unique values
output: a new array containing all values held in common
between the inputs

e.g. [1, 2, 3, 4, 5] and [0, 2, 4, 6, 8] return [2, 4]

rules: must be O(N) or less

determine shortest argument array and set it as one to be hashed

create a an "index" hash table from one array
- this means a hash where the keys are array elements
and the values are a truthy value
- which array? the shortest array is preferable to hash to reduce memory requirements

create an empty array
- worst case, empty result array should be same size as smallest array

loop though array that wasn't "indexed"
- if current element found in hash table,
add current element to result array

return result array
*/

package main

import (
	"fmt"
)

func intArrayIntersection(sl1, sl2 []int) []int {
	shortest := sl1
	longest := sl2

	if len(sl2) < len(sl1) {
		shortest = sl2
		longest = sl1
	}

	index := make(map[int]bool, len(shortest))
	for _, elem := range shortest {
		index[elem] = true
	}

	result := []int{}
	for _, elem := range longest {
		if index[elem] {
			result = append(result, elem)
		}
	}

	return result
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{0, 2, 4, 6, 8}

	slice3 := intArrayIntersection(slice1, slice2)
	fmt.Println(slice3)
}
