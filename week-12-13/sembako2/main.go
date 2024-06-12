package main

import "fmt"

type barang struct {
	nama               string
	harga, panjangNama int
}

const NMAX int = 1024

type dataBarang [NMAX]barang

func isiArray(himpunan *dataBarang, n int) {
	//algoritma untuk menginput data himpunan barang ke dalam bentuk array
	var nama string
	var harga int

	for i := 0; i < n; i++ {
		// gunakan fmt.Scan untuk menginput nilai di sini
		fmt.Scanln(&nama, &harga)
		// proses menghimpun nilai sesuai jumlah n
		himpunan[i].nama = nama
		himpunan[i].harga = harga
		himpunan[i].panjangNama = len(nama)
	}
}

func insertionSort(himpunan *dataBarang, n int) {
	//algoritma insertion sort secara ascending
	var pass, i int
	var temp barang

	pass = 1

	for pass <= n-1 {
		i = pass
		temp = himpunan[pass]
		for i > 0 && temp.panjangNama < himpunan[i-1].panjangNama {
			himpunan[i] = himpunan[i-1]
			i--
		}
		himpunan[i] = temp
		pass++
	}
}

func showArray(himpunan dataBarang, n int) {
	//algoritma untuk menampilkan data barang menggunakan pengulangan array
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
