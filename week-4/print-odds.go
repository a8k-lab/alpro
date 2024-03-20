package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	printFactorNumbers(n)
}

func printOddNumbers(n int) {
	if n <= 0 {
		return
	}

	printFactorNumbers(n - 1)

	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}
