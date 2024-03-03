package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)
	fmt.Println(sumOdd(num))
}

func sumOdd(num int) int {
	// returing 1 + 3 + 5 + ... N
	result := 0

	for i := 1; i <= num; i += 2 {
		result += i
	}

	return result
}
