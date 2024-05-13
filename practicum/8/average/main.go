package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var A tabInt
	var n int
	fmt.Scan(&n)
	baca(&A, n)
	fmt.Println(rataRata(A, n))
	fmt.Println(lebihDariRataRata(A, n))
}

func baca(A *tabInt, n int) {
	/*
		IS: A terdefinisi sembarang, n terdefinisi. Asumsi n <= NMAX
		FS: A berisi nilai
	*/
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func rataRata(A tabInt, n int) float64 {
	/* Mengembalikan rata-rata nilai elemen A */
	var sum int = 0
	for i := 0; i < n; i++ {
		sum += A[i]
	}

	return float64(sum) / float64(n)
}

func lebihDariRataRata(A tabInt, n int) int {
	/* Menghitung banyak elemen array yang nilainya lebih dari nilai rata-rata */
	var total int = 0
	for i := 0; i < n; i++ {
		// jika A[i] lebih besar dari rata-rata, maka total di-increment
		if A[i] > int(rataRata(A, n)) {
			total++
		}
	}

	return total
}
