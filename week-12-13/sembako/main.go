package main

import "fmt"

type barang struct {
	nama  string
	harga int
}

const NMAX int = 1024

type dataBarang [NMAX]barang

func isiArray(himpunan *dataBarang, n int) {
	// algoritma untuk menginput data barang dalam bentuk array
	// hint : gunakan variabel lokal dan input scan
	var nama string
	var harga int

	for i := 0; i < n; i++ {
		fmt.Scanln(&nama, &harga)
		himpunan[i].nama = nama
		himpunan[i].harga = harga
	}
}

func insertionSort(himpunan *dataBarang, n int) {
	//algoritma insertion sort secara descending
	var pass, i int
	var temp barang

	pass = 1

	for pass <= n-1 {
		i = pass
		temp = himpunan[pass]
		for i > 0 && temp.harga > himpunan[i-1].harga {
			himpunan[i] = himpunan[i-1]
			i--
		}
		himpunan[i] = temp
		pass++
	}
}

func showArray(himpunan dataBarang, n int) {
	//algoritma menampilkan data barang
	for i := 0; i < n; i++ {
		fmt.Println(himpunan[i].nama, himpunan[i].harga)
	}
}

func main() {
	var himpunan dataBarang
	var n int
	fmt.Scanln(&n)
	isiArray(&himpunan, n)
	insertionSort(&himpunan, n)
	showArray(himpunan, n)
}
