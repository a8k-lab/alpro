package main

import "fmt"

func median(arrayN []int, n int) float64 {
	if n%2 == 0 {
		fmt.Println(arrayN[n/2-1], arrayN[n/2])
		return float64(arrayN[n/2-1]+arrayN[n/2]) / 2
	}

	return float64(arrayN[n/2])
}

func main() {
	var n int

	fmt.Println("Masukkan jumlah bilangan")
	fmt.Scan(&n)

	arrayN := make([]int, n)
	fmt.Println("Masukkan bilangan")
	for i := 0; i < n; i++ {
		fmt.Scan(&arrayN[i])
	}

	fmt.Println("Median adalah", median(arrayN, n))
}
