package main

import "fmt"

const N int = 4

func searchBil(T [N]int, x, n int) int {
	for i := 0; i < n; i++ {
		if T[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	var arr = [N]int{1, 2, 3, 4}

	fmt.Println(searchBil(arr, 10, 4))
}
