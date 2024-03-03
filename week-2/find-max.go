package main

import "fmt"

func main() {
	var a, b, c, biggest int
	findMax(a, b, c, biggest)
}

func findMax(a, b, c, biggest int) {
	fmt.Scan(&a, &b, &c)
	biggest = a
	if b > a && b > c {
		biggest = b
	}
	if c > a && c > b {
		biggest = c
	}
	fmt.Println(biggest)
}
