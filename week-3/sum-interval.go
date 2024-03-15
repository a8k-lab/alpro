package main

import "fmt"

func main() {
	var start, end int

	fmt.Scan(&start, &end)
	sumInterval(end, start)
}

func sumInterval(M, N int) {
	sum := 0
	for i := M; i >= N; i-- {
		sum += i
	}

	fmt.Println(sum)
}
