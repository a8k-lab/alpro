package main

import "fmt"

const NMAX int = 5

type matrix [NMAX][NMAX]int

func main() {
	var m1, m2, m3 matrix
	var n int

	fmt.Scan(&n) // ukuran matriks
	baca(&m1, &n)
	baca(&m2, &n)

	fmt.Println("Matriks pertama")
	cetak(m1, n)

	fmt.Println("Matriks kedua")
	cetak(m2, n)

	fmt.Println("Matriks ketiga")
	jumlah(m1, m2, &m3, n)
	cetak(m3, n)
}

func baca(m *matrix, n *int) {
	/*
		IS: Matriks m terdefinisi sembarang, n terdefinisi
		FS: Matriks m dengan baris n x kolom n berisi nilai. Jika n > NMAX
		    gunakan nilai n = NMAX
	*/
	var inputN int
	// baris dan kolom digunakan untuk menentukan baris dan kolom matriks
	// baris = baris, kolom = kolom
	var baris, kolom int = 0, 0

	for i := 0; i < *n**n; i++ {
		fmt.Scan(&inputN)

		if i < NMAX*NMAX {
			// masukkan input ke array `m` hanya jika ketika input masih berada sebelum ke n*n
			m[baris][kolom] = inputN
			kolom++

			// jika kolom sudah habis
			// maka kembali ke kolom pertama dan lanjut ke baris selanjutnya
			if kolom == NMAX {
				baris++
				kolom = 0
			}
		}
	}
}

func cetak(m matrix, n int) {
	/*
		IS: Matrik m terdefinisi, n terdefinisi
		FS: Tercetak matriks m di layar dengan format:

			x11 x12 ... x1n
			x21 x22 ... x2n
			... ... ... ...
			xn1 xn2 ...	xnn

			dengan n adalah ukuran matriks
	*/
	for baris := 0; baris < n; baris++ {
		for kolom := 0; kolom < n; kolom++ {
			if baris < NMAX && kolom < NMAX {
				fmt.Printf("%d ", m[baris][kolom])
			}
		}
		fmt.Println()
	}
}

func jumlah(A, B matrix, C *matrix, n int) {
	/*
		IS: Matriks A dan B terdefinisi
		FS: Matriks C berisi nilai. Elemen ke-baris dan j matriks C merupakan
		    penjumlahan elemen ke-baris dan j dari matriks A dan B.

			Contoh matriks dengan n = 2:
			matriks A + matriks B = matriks C
			1 1		 +   3 0	  =  4 1
			2 2			 4 1		 6 3
	*/
	for baris := 0; baris < n; baris++ {
		for kolom := 0; kolom < n; kolom++ {
			if baris < NMAX && kolom < NMAX {
				C[baris][kolom] = A[baris][kolom] + B[baris][kolom]
			}
		}
	}
}
