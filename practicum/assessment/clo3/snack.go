package main

import "fmt"

// Konstanta untuk ukuran maksimum elemen array sebanyak 10
const NMAX = 10

// Tipe bentukan struktur untuk menampung data snack
type snack struct {
	id    string
	nama  string
	harga int
	stok  int
}

// tipe alias array of snack
type tabSnack [NMAX]snack

func main() {
	// Deklarasi variabel
	var data tabSnack
	var nData int
	var id string
	var harga int

	// Baca data
	bacaData(&data, &nData)

	// cetak data awal
	cetakData(data, nData)

	// cari nama snack berdasarkan id
	fmt.Scan(&id)
	fmt.Println("Nama snack:", namaSnackSeq(data, nData, id))

	// Hitung dan cetak omset
	fmt.Println("Omset:", omset(data, nData))

	// Mengurutkan array berdasarkan ID secara menaik
	urutID(&data, nData)

	// cetak data setelah diurutkan
	cetakData(data, nData)

	// Mengurutkan array berdasarkan harga secara menurun
	urutHarga(&data, nData)

	// cetak data setelah diurutkan
	cetakData(data, nData)

	// cari nama snack berdasarkan harga
	fmt.Scan(&harga)
	fmt.Println("Nama snack:", namaSnackBin(data, nData, harga))
}

func bacaData(A *tabSnack, n *int) {
	/*
		IS: Array A terdefinisi sembarang
		Proses : Selama array belum penuh dan yang dibaca bukan "none" "none" 0 0,
				 maka data dimasukkan ke dalam array A
		FS: Array A terisi sebanyak n elemen
	*/
	var snackInput snack
	var i int = 0

	for {
		fmt.Scan(&snackInput.id)
		fmt.Scan(&snackInput.nama)
		fmt.Scan(&snackInput.harga)
		fmt.Scan(&snackInput.stok)

		if snackInput.id == "none" && snackInput.nama == "none" && snackInput.harga == 0 && snackInput.stok == 0 {
			return
		} else if *n < NMAX {
			A[i] = snackInput
			*n++
		}
		i++
	}
}

func cetakData(A tabSnack, n int) {
	/*
		IS: Terdefinisi Array A  sebanyak n elemen
		FS: Tercetak di layar seluruh atribut A dengan format seperti contoh
	*/
	var i int
	fmt.Println("=====================================")
	fmt.Println("             Data Snack         ")
	fmt.Println("=====================================")

	for i < n {
		fmt.Printf("%5s", A[i].id)
		fmt.Printf("%18s", A[i].nama)
		fmt.Printf("%6d", A[i].harga)
		fmt.Printf("%6d\n", A[i].stok)
		i++
	}

	fmt.Println("=====================================")
}

func namaSnackSeq(A tabSnack, n int, x string) string {
	/* Diberikan array A sebanyak n elemen terurut acak, fungsi mengembalikan
	   nama snack dengan id (x) jika x terdapat pada array A. Jika tidak
	   terdapat pada array A, maka kembalikan string "None" */
	for i := 0; i < n; i++ {
		if A[i].id == x {
			// jika ketemu, keluar dari perulangan dan return nama ke fungsi
			return A[i].nama
		}
	}

	return "None"
}

func omset(A tabSnack, n int) int {
	/*
		Diberikan array A sebanyak n elemen, fungsi mengembalikan omset
		yang merupakan jumlah total dari perkalian antara stok dengan harga
	*/
	var total int = 0

	for i := 0; i < n; i++ {
		total += A[i].stok * A[i].harga
	}

	return total
}

func urutID(A *tabSnack, n int) {
	/*
		IS: Terdefinisi array A sebanyak n elemen dengan terurut sembarang
		FS: Urutkan array A secara menaik (ascending) berdasarkan ID dengan
			algoritma selection sort
	*/
	var pass, idx, i int
	var temp snack

	pass = 1

	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[idx].id > A[i].id {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func urutHarga(A *tabSnack, n int) {
	/*
		IS: Terdefinisi array A sebanyak n elemen dengan terurut sembarang
		FS: Urutkan array A secara menurun (descending) berdasarkan harga dengan
		    algoritma insertion sort
	*/
	var pass, i int
	var temp snack

	pass = 1

	for pass <= n-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.harga < A[i-1].harga {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

func namaSnackBin(A tabSnack, n int, x int) string {
	/* Diberikan array A sebanyak n elemen terurut menurun berdasarkan harga,
	   fungsi mengembalikan nama snack dengan harga (x) jika x terdapat pada
	   array A. Jika tidak terdapat pada array A, maka kembalikan string "None"
	*/
	var found, left, mid, right int

	left = 1
	right = n - 1
	found = -1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x < A[mid].harga {
			// pindahkan pointer right ke tengah
			right = mid - 1
		} else if x > A[mid].harga {
			// pindahkan pointer kiri ke tengah
			left = mid + 1
		} else {
			// selain 2 kondisi di atas, maka bisa dipastikan x berada di tengah
			found = mid
		}
	}

	if found == -1 {
		return "None"
	}
	return A[found].nama
}
