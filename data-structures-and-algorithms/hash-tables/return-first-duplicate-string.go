/*
input: array of strings
output: first duplicate string (first string that appears twice) in array

assume the array has at least one pair of duplicates
must be O(N)

create an empty hash table
for every element in the array
	- if element is in hash table,
		return from function with element
	- else
		add to hash table as table[elem] = true
return "no duplicates"
*/

package main

import (
	"fmt"
)

func duplicateString(slice []string) string {
	table := make(map[string]bool)
	for _, elem := range slice {
		if table[elem] {
			return elem
		} else {
			table[elem] = true
		}
	}
	return "no duplicates"
}

func main() {
	slice := []string{"a", "b", "c", "d", "c", "e", "f"}
	fmt.Println(duplicateString(slice))
}
