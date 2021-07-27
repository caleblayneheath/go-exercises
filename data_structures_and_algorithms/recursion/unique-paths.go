package main

import (
	"fmt"
)

func uniquePaths(rows, cols int) int {
	if (rows <= 1) || (cols <= 1) {
		return 1
	}

	return uniquePaths(rows-1, cols) + uniquePaths(rows, cols-1)

}

func main() {
	fmt.Println(uniquePaths(3, 3))
}

