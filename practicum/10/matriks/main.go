package main

import "fmt"

const ROWMAX int = 5
const COLUMNMAX int = 5

type tabInt [ROWMAX][COLUMNMAX]int

func main() {
	var matriks tabInt
	var nRow, nCol int // nRow adalah banyak elemen baris, nCol adalah banyak elemen baris

	// membaca banyak elemen baris dan elemen kolom
	fmt.Scan(&nRow, &nCol)

	// membaca matriks nRow x nCol
	baca(&matriks, nRow, nCol)

	// cetak matriks nRow x nCol
	cetak(matriks, nRow, nCol)

	// cetak nilai minimum matriks dengan posisi baris dan kolomnya
	minimum(matriks, nRow, nCol)
}

func baca(A *tabInt, nR int, nC int) {
	/*
		IS: Matriks A terdefinisi sembarang. Banyak elemen baris dan kolom (nR dan nC) terdefinisi
		FS: Matriks A nR x nC berisi nilai
	*/
	// iR = iterationRow, iC = iterationColumn :3
	for iR := 0; iR < nR; iR++ {
		for iC := 0; iC < nC; iC++ {
			fmt.Scan(&A[iR][iC])
		}
	}
}

func cetak(A tabInt, nR int, nC int) {
	/*
			IS: Matriks A nR x nC terdefinisi
			FS: Matriks A nR x nC tercetak di layar dengan format:
				"Data matriks:
				 <e11> <e12> <e13> ...
		 		 <e21> <e22> <e23> ...
				 ...
				 ...	...	  ...  <enRnC>"
	*/
	fmt.Println("Data matriks:")
	// iR = iterationRow, iC = iterationColumn :3
	for iR := 0; iR < nR; iR++ {
		for iC := 0; iC < nC; iC++ {
			fmt.Printf("%d ", A[iR][iC])
		}
		fmt.Println()
	}
}

func minimum(A tabInt, nR int, nC int) {
	/*
		IS: Matriks A nRxnC terdefinisi
		FS: Tercetak di layar nilai minimum matriks dengan posisi baris dan kolomnya dengan format:
			"Nilai minimum adalah <min> pada posisi baris <baris> dan kolom <kolom>"
			Catatan: Baris dan kolom dimulai dari angka 1
	*/
	// inisialisasi awal nilai minimum menjadi A[0][0] atau elemen matriks baris 1 kolom 1
	var min int = A[0][0]
	// minR = minimumRow, minC = minimumColumn
	var minR int = 1
	var minC int = 1

	// iR = iterationRow, iC = iterationColumn :3
	for iR := 0; iR < nR; iR++ {
		for iC := 0; iC < nC; iC++ {
			if A[iR][iC] < min {
				min = A[iR][iC]
				// ditambahkan 1 karena iR dan iC adalah nilai index yang dimulai dari 0
				// sedangkan minR dan minC akan berisi posisi matriks yang nilainya dimulai dari 1
				minR = iR + 1
				minC = iC + 1
			}
		}
	}

	fmt.Printf("Nilai minimum adalah %d pada posisi baris %d kolom %d\n", min, minR, minC)
}
