package main

import "fmt"

func main() {
	var rows int

	fmt.Scan(&rows)
	starRecursive(rows)
}

func starRecursive(rows int) {
	if rows <= 0 {
		return
	}

	starRecursive(rows - 1)
	for i := 0; i < rows; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
