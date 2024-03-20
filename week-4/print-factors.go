package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	printFactorNumbers(n)
}

func printFactorNumbers(n int) {
	if n <= 0 {
		return
	}

	printFactorNumbers(n - 1)

	if n%n != 0 {
		fmt.Print(n, " ")
	}
}
