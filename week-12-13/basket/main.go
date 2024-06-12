package basket

import "fmt"

type pemain struct {
	nama              string
	poin, panjangNama int
}

const NMAX int = 1024

type dataPemain [NMAX]pemain

func isiArray(himpunan *dataPemain, n int) {
	//algoritma untuk menginput data pemain dalam bentuk array
	// gunakan variabel lokal dan input scan
	var nama string
	var poin int

	for i := 0; i < n; i++ {
		fmt.Scanln(&nama, &poin)
		himpunan[i].nama = nama
		himpunan[i].poin = poin
		himpunan[i].panjangNama = len(nama)
	}
}

func selectionSort(himpunan *dataPemain, n int) {
	//algoritma selection sort secara descending
	var pass, idx, i int
	var temp pemain

	pass = 1

	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if himpunan[idx].panjangNama < himpunan[i].panjangNama {
				idx = i
			}
			i++
		}
		temp = himpunan[pass-1]
		himpunan[pass-1] = himpunan[idx]
		himpunan[idx] = temp
		pass++
	}
}

func showArray(himpunan dataPemain, n int) {
	//algoritma untuk menampilkan data array
	for i := 0; i < n; i++ {
		fmt.Println(himpunan[i].nama, himpunan[i].poin)
	}
}

func main() {
	var himpunan dataPemain
	var n int
	fmt.Scanln(&n)

	isiArray(&himpunan, n)
	selectionSort(&himpunan, n)
	showArray(himpunan, n)
}
