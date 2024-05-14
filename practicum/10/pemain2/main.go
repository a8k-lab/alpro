package main

import "fmt"

const NMAX int = 24

// Deklarasi tipe bentukan struktur pemain
type pemain struct {
	nama, noPunggung, posisi string
	tinggi                   int
}

// deklarasi tipe alias tabPemain dengan maksimal NMAX elemen
type tabPemain [NMAX]pemain

func main() {
	var timnas tabPemain
	var nPemain int
	var p string // p adalah posisi pemain

	// baca banyaknya pemain yang akan diinputkan
	fmt.Scan(&nPemain)

	// baca data pemain
	bacaData(&timnas, &nPemain)

	// baca posisi pemain (p)
	fmt.Scan(&p)

	// cetak data pemain
	cetakPemain(timnas, nPemain)

	// cetak nama pemain tertinggi pada posisi (p)
	fmt.Printf("Pemain tertinggi pada posisi %s: %s\n", p, pemainTertinggi(timnas, nPemain, p))
}

func bacaData(A *tabPemain, n *int) {
	/*
		IS: Array A terdefinisi sembarang. Banyak elemen n terdefinisi.
		FS: Array A sebanyak n elemen berisi nilai. Jika n > NMAX, maka n = NMAX
	*/
	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i].nama)
		fmt.Scan(&A[i].noPunggung)
		fmt.Scan(&A[i].posisi)
		fmt.Scan(&A[i].tinggi)
	}
}

func cetakPemain(A tabPemain, n int) {
	/*
		IS: Array A yang bertipe tabPemain dengan banyak elemen n terdefinisi
		FS: Tercetak di layar elemen array A sebanyak n dengan format:
			"Data pemain:
			<nama1> <nomorPunggung1> <posisi1> <tinggi1>
			<nama2> <nomorPunggung2> <posisi2> <tinggi2>
			...
			<naman> <nomorPunggungn> <posisin> <tinggin>"
	*/
	fmt.Println("Data Pemain:")

	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %s %d\n", A[i].nama, A[i].noPunggung, A[i].posisi, A[i].tinggi)
	}
}

func pemainTertinggi(A tabPemain, n int, pos string) string {
	/* Diberikan array A bertipe tabPemain dengan banyak elemen n dan posisi (pos), fungsi akan
	   mengembalikan nama pemain yang memiliki tinggi badan tertinggi pada posisi (pos) itu */
	var tertinggi pemain

	for i := 0; i < n; i++ {
		if A[i].posisi == pos { // jika pemain memiliki posisi yang sama dengan argument `pos`
			if A[i].tinggi > tertinggi.tinggi {
				tertinggi = A[i]
			}
		}
	}

	return tertinggi.nama
}
