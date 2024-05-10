package main

import "fmt"

const NMAX int = 5

type tabInt struct {
	info [NMAX]int
	n    int
}

func main() {
	var nilaiAkhir tabInt

	bacaNilai(&nilaiAkhir)
	cetakNilai(nilaiAkhir)
}

func bacaNilai(NA *tabInt) {
	for i := 0; i < NMAX; i++ {
		fmt.Scan(&NA.info[i])
		NA.n++
	}
}

func cetakNilai(NA tabInt) {
	for i := 0; i < NA.n; i++ {
		fmt.Printf("%d ", NA.info[i])
	}
}
