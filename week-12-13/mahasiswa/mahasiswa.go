package mahasiswa

import "fmt"

type mahasiswa struct {
	nama, indeks string
	nilai        int
}

const NMAX int = 1024

type dataMahasiswa [NMAX]mahasiswa

func indeksMahasiswa(nsm int) string {
	if nsm > 80 {
		return "A"
	} else if nsm > 70 {
		return "AB"
	} else if nsm > 65 {
		return "B"
	} else if nsm > 60 {
		return "BC"
	} else if nsm > 50 {
		return "C"
	} else if nsm > 40 {
		return "D"
	}
	return "E"
}

func isiArray(himpunan *dataMahasiswa, n int) {
	//algoritma untuk menginput data mahasiswa dalam bentuk array dan index nilai
	// note :  gunakan if else untuk menampilkan nilai oindex
	var nama string
	var nilai int

	for i := 0; i < n; i++ {
		fmt.Scanln(&nama, &nilai)
		himpunan[i].nama = nama
		himpunan[i].nilai = nilai
		himpunan[i].indeks = indeksMahasiswa(nilai)
	}
}

func insertionSort(himpunan *dataMahasiswa, n int) {
	//algoritma insertion sort secara ascending
	var pass, i int
	var temp mahasiswa

	for pass <= n-1 {
		i = pass
		temp = himpunan[pass]
		for i > 0 && temp.nilai < himpunan[i-1].nilai {
			himpunan[i] = himpunan[i-1]
			i--
		}
		himpunan[i] = temp
		pass++
	}
}

func showArray(himpunan dataMahasiswa, n int, tampilIndeks string) {
	//algoritma untuk menampilkan data menggunakan pengulangan array
	for i := 0; i < n; i++ {
		if himpunan[i].indeks == tampilIndeks {
			fmt.Println(himpunan[i].nama, himpunan[i].nilai, himpunan[i].indeks)
		}
	}
}

func main() {
	var himpunan dataMahasiswa
	var n int
	var tampilIndeks string
	fmt.Scanln(&n)
	fmt.Scanln(&tampilIndeks)
	isiArray(&himpunan, n)
	insertionSort(&himpunan, n)
	showArray(himpunan, n, tampilIndeks)
}
