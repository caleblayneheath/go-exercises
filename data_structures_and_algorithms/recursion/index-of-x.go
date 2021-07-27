package main

import (
	"fmt"
)

func indexOfx(str string, index ...int) int {
	// return -1 if no "x" present in string

	// guard clause for empty string
	if len(str) == 0 {
		return -1
	}

	// default value for index parameter is 0
	var idx int
	if len(index) == 0 {
		idx = 0
	} else {
		idx = index[0]
	}

	// current char is "x",
	// call recursively if more chars to examine,
	// if none left return -1 for none found
	if string(str[idx]) == "x" {
		return idx
	} else if len(str) > (idx + 1) {
		return indexOfx(str, idx+1)
	} else {
		return -1
	}

}

func main() {

	str1 := "abcdefghijklmnopqrstuvwxyz"
	str2 := "x"
	str3 := "12xx"
	str4 := "abcd"
	str5 := ""

	fmt.Println(indexOfx(str1)) // 23
	fmt.Println(indexOfx(str2)) // 0
	fmt.Println(indexOfx(str3)) // 2
	fmt.Println(indexOfx(str4)) // -1
	fmt.Println(indexOfx(str5)) // -1
	fmt.Println("")

}

