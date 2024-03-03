package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)
	fmt.Println(isPrime(num))
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
