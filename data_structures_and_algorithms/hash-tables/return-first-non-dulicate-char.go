/*
input: string
output: first char that doesn't appear more than once, or else ""

rules:
must be O(N)

to know there is no duplicate, you must traverse entire string at least once
you must ascertain position of the character

keep letter count table
iterate through entire string
- if char in table
	table[char] += 1
-else
	table[char = 1]

iterate through string again
- if table[char] > 1
	continue
- else
	return char

*/

package main

import (
	"fmt"
)

func nonDuplicateChar(str string) string {
	count := map[rune]int{}

	for _, c := range str {
		count[c] += 1
	}

	for _, c := range str {
		if count[c] == 1 {
			return string(c)
		}
	}
	return ""
}

func main() {
	str := "minimum"
	fmt.Println(nonDuplicateChar(str)) // print 'n'
}
