package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Scan(&a, &b, &n)
	fmt.Println(unarithmathic(a, b, n))
}

func unarithmathic(a, b, n int) int {
	if n == 1 {
		return a
	} else {
		return unarithmathic(a+b, b, n-1)
	}
}
