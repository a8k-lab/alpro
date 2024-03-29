package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	fmt.Printf("%.2f", numSeries(n))
}

func numSeries(n int) float64 {
	if n == 1 {
		return 1
	} else {

		return 1.0/float64(n) + numSeries(n-1)
	}
}
