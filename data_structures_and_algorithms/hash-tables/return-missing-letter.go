/*
input: string containing at minimum all alphabetic chars but one
output: string that is the missing alphabet char

rules: must be O(N)
case doesn't matter

create hash table
iterate through string
- if char is alphabetical,
	table[char.toUpper] = true
iterate through alphabet string
- return char if table[char] != char
*/

package main

import (
	"fmt"
	"unicode"
)

func returnMissingLetter(str string) string {
	letters := make(map[rune]bool, 25)

	for _, c := range str {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			letters[unicode.ToUpper(c)] = true
		}
	}

	result := ""
	for _, c := range "ABCDEFGHHIJKLMNOPQRSTUVWXYZ" {
		if !letters[c] {
			result = string(c)
			break
		}
	}
	return result
}

func main() {
	sentence := "the quick brown box jumps over a lazy dog"
	fmt.Println(returnMissingLetter(sentence))
}
