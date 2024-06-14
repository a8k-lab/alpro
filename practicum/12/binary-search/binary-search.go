package main

import "fmt"

const NMAX = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData, x2, idx int

	fmt.Scan(&x2)

	bacaData(&data, &nData)

	cetakData(data, nData)

	fmt.Print("Hasil pencarian: ")
	binarySearch(data, nData, x2, &idx)

	if idx != -1 {
		fmt.Printf("Bilangan %d ditemukan pada posisi ke-%d\n", x2, idx+1)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	var i int = 0
	*n = 0

	for {
		if i == NMAX {
			return
		}

		fmt.Scan(&A[i])

		if i > 0 {
			if A[i-1] > A[i] {
				return
			}
		}
		i++
		*n++
	}
}

func cetakData(A tabInt, n int) {
	fmt.Print("Data Bilangan: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func binarySearch(A tabInt, n, x int, idx *int) {
	var left, found, mid, right int

	left = 1
	right = n - 1
	found = -1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x < A[mid] {
			right = mid - 1
		} else if x > A[mid] {
			left = mid + 1
		} else {
			found = mid
		}
	}

	*idx = found
}
