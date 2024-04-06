package main

import "fmt"

func lucas(n int) int {
	if n == 1 {
		return 2
	} else if n == 2 {
		return 1
	}

	return lucas(n-2) + lucas(n-1)
}

func main() {
	var n int

	fmt.Scan(&n)
	fmt.Println(lucas(n))
}
