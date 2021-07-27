package main

import (
	"fmt"
)

func characterCount(slice []string) int {
	if len(slice) == 0 {
		return 0
	}
	
	if len(slice) == 1 {
		return len(slice[0])
	} else {
		return len(slice[0]) + characterCount(slice[1:])
	}

}

func main() {
	arr := []string{"ab", "c", "def", "ghij"}
	fmt.Println(characterCount(arr)) // 10
	fmt.Println(characterCount([]string{})) // 0

}

