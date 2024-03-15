package main

import "fmt"

func main() {
	var a, b int
	isLoop := true

	for isLoop {
		fmt.Scan(&a, &b)
		if a == b {
			isLoop = false
		} else {
			sortTwoInt(&a, &b)
			fmt.Println(a, b)
		}
	}
}

func sortTwoInt(first, second *int) {
	if *first > *second {
		*first = *first + *second
		*second = *first - *second
		*first = *first - *second
	}
}
