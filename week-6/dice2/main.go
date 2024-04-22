package main

import "fmt"

type SkorGame struct {
	TotalA   int
	TotalB   int
	TotalC   int
	min      int
	Pemenang byte
}

var skor SkorGame

func minSkor() {
	skor.min = skor.TotalA
	skor.Pemenang = 65 // 65 = nilai byte untuk "A"
	if skor.TotalB < skor.TotalA && skor.TotalB < skor.TotalC {
		skor.min = skor.TotalB
		skor.Pemenang = 66 // 66 = nilai byte untuk "B"
	} else if skor.TotalC < skor.TotalA && skor.TotalC < skor.TotalB {
		skor.min = skor.TotalC
		skor.Pemenang = 67 // 67 = nilai byte untuk "C"
	}
}

func gameDadu(n int) {
	var d1, d2, d3 int

	for i := 0; i < n; i++ {
		fmt.Scan(&d1, &d2, &d3)
		skor.TotalA += d1
		skor.TotalB += d2
		skor.TotalC += d3
	}
	minSkor()
}

func main() {
	var n int

	fmt.Scan(&n)
	gameDadu(n)
	fmt.Printf("%c %d\n", skor.Pemenang, skor.min)
}
