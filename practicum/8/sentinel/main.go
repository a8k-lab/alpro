package main

import (
	"fmt"
	"math"
)

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int

	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), rata_rata(data, nData))
}

func baca(A *tabInt, n *int) {
	var isLoop bool = true
	var i int = 0
	*n = 0

	for isLoop {
		if *n < NMAX {
			fmt.Scan(&A[i])

			if A[i] == 0 {
				isLoop = false
			} else {
				*n++
				i++
			}
		} else {
			isLoop = false
		}
	}
}

func cetak(A tabInt, n int) {
	for i := 0; i < n; i++ {
		if A[i] > 0 {
			fmt.Printf("%d ", A[i])
		}
	}
	fmt.Println()
}

func jumlah(A tabInt, n int) int {
	var sum int = 0
	for i := 0; i < n; i++ {
		// jika angka merupakan negatif
		// maka diubah menjadi positif dahulu sebelum ditambah ke `sum`
		if A[i] < 0 {
			sum += A[i] * -1
		} else {
			sum += A[i]
		}
	}

	return sum
}

func rata_rata(A tabInt, n int) float64 {
	// jika n = 0
	// maka kembalikan NaN (Not a Number) yang diambil dari package `math` bawaan Go
	if n == 0 {
		return math.NaN()
	}
	return float64(jumlah(A, n)) / float64(n)
}
