package main

import "fmt"

func main() {
	var sum, start, end int

	sum = 0
	fmt.Scan(&start, &end)

	for start <= end {
		sum += start
		start++
	}

	fmt.Println(sum)
}
