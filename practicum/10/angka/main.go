package main

import "fmt"

const NMAX int = 20

type tabInt [NMAX]int

func main() {
	var arrAngka tabInt
	var panjangArr int

	baca(&arrAngka, &panjangArr)
	cetakElemen(arrAngka, panjangArr)
	cetakInfo(arrAngka, panjangArr)
}

func baca(A *tabInt, n *int) {
	var isLoop bool = true
	var i int = 0
	*n = 0

	for isLoop {
		if *n < NMAX {
			fmt.Scan(&(*A)[i])

			if (*A)[i] < 0 {
				isLoop = false
			} else if (*A)[i] != 0 {
				*n++
				i++
			}
		} else {
			isLoop = false
		}
	}
}

func cetakElemen(A tabInt, n int) {
	fmt.Print("Elemen array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func maksimum(A tabInt, n int) int {
	var maks int = A[0]
	for i := 0; i < n; i++ {
		if A[i] > maks {
			maks = A[i]
		}
	}

	return maks
}

func minimum(A tabInt, n int) int {
	var min int = A[0]
	for i := 0; i < n; i++ {
		if A[i] < min {
			min = A[i]
		}
	}

	return min
}

func cetakInfo(A tabInt, n int) {
	fmt.Printf("Nilai maksimum: %d\n", maksimum(A, n))
	fmt.Printf("Nilai minimum: %d\n", minimum(A, n))
	fmt.Printf("Banyak elemen: %d\n", n)
}
