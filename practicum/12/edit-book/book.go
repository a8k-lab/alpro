package main

import "fmt"

// Deklarasi konstanta global NMAX sebanyak 10
const NMAX = 10

/*
Deklarasi tipe bentukan struktur books dengan atribut:

	bookId (string), tittle (string), author (string), genre (string)
	pubYear (int)
*/
type book struct {
	bookId  string
	title   string
	author  string
	genre   string
	pubYear int
}

// Deklarasi tibe alias tabBook yang merupakan array of books dengan dimensi NMAX
type tabBook [NMAX]book

func main() {
	// Deklarasi variabel lokal
	var data tabBook
	var nData int
	var bookId, genre string

	// Pembacaan data
	fmt.Scan(&nData)
	readData(&data, &nData)

	// Pembacaan bookId yang mau diedit dan genre penggantinya
	fmt.Scan(&bookId, &genre)

	// Data buku sebelum diedit
	printData(data, nData)

	// Prosedur pengeditan
	editData(&data, nData, bookId, genre)

	// Data buku setelah diedit
	printData(data, nData)
}

func readData(A *tabBook, n *int) {
	/*
		IS: Array A sebanyak n elemen terdefinisi sembarang
		Proses: 1. Membaca banyak elemen (n)
				2. Jika n > NMAX, n = NMAX
				3. Baca seluruh atribut A dan masukkan ke array sebanyak n elemen
		FS: Array A sebanyak n elemen berisi data.
	*/
	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i].bookId)
		fmt.Scan(&A[i].title)
		fmt.Scan(&A[i].author)
		fmt.Scan(&A[i].genre)
		fmt.Scan(&A[i].pubYear)
	}
}

func printData(A tabBook, n int) {
	/*
		IS: Array A sebanyak n elemen terdefinisi
		FS: Tercetak di layar data buku dengan format:
			"Books Data:
			 <bookId1> <title1> <author1> <genre1> <pubYear1>
			 ...		...		...		  ...		...
			 <bookIdn> <titlen> <authorn> <genren> <pubYearn>"
	*/
	fmt.Println("Books Data:")

	for i := 0; i < n; i++ {
		fmt.Printf("%s ", A[i].bookId)
		fmt.Printf("%s ", A[i].title)
		fmt.Printf("%s ", A[i].author)
		fmt.Printf("%s ", A[i].genre)
		fmt.Printf("%d\n", A[i].pubYear)
	}
}

func editData(A *tabBook, n int, bId string, x string) {
	/*
		IS: Array A sebanyak n elemen terdefinisi.
			bookId yang akan diedit (bId) terdefinisi
			genre (x) terdefinisi
		Proses: 1. Lakukan pencarian idx buku berdasarkan bookId (bId)
				   dengan menggunakan fungsi findData.
				2. a. Jika bId buku yang akan diedit berhasil ditemukan, maka
				   lakukan pengubahan genre dengan nilai x.
				   Lalu cetak di layar "Data has been edited"
				   b. Jika bId buku yang akan diedit tidak berhasil ditemukan,
				   maka cetak di layar "Data has not been edited".
		FS: Array A sebanyak n elemen berubah isinya, jika bId buku ada
			dalam array (berhasil ditemukan). Jika tidak ada, maka Array tidak
			berubah.
	*/
	// 1.
	var idx int = findData(*A, n, bId)

	// 2.
	if idx != -1 {
		A[idx].genre = x
		fmt.Println("Data has been edited")
	} else {
		fmt.Println("Data has not been edited")
	}
}

func findData(A tabBook, n int, bId string) int {
	/* Diberikan array A sebanyak n elemen dan book Id (bId) dari buku, fungsi
	   akan mengembalikan indeks buku itu jika ditemukan, dan mengembalikan
	   -1 jika tidak ditemukan. Catatan: Gunakan algoritma sequential search */
	for i := 0; i < n; i++ {
		if A[i].bookId == bId {
			return i
		}
	}

	return -1
}
