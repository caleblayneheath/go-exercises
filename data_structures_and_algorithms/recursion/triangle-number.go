package main

import (
	"fmt"
)

func triangleNumber(n int) int {
	// assume n <= 1 should return 1

	if n <= 1 {
		return 1
	} else {
		return n + triangleNumber(n-1)
	}

}

func main() {
	fmt.Println(triangleNumber(7) == 28)
	fmt.Println(triangleNumber(0) == 1)
	fmt.Println(triangleNumber(1) == 1)
	fmt.Println(triangleNumber(5) == 15)
}

