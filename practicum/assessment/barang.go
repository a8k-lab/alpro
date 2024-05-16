package main

import "fmt"

// Deklarasi tipe tBarang dengan field: id (string), nama (string), harga (integer), stok (integer)
type tBarang struct {
	id, nama    string
	harga, stok int
}

// Deklarasi konstanta NMAX dengan nilai integer 10
const NMAX int = 10

// Deklarasi tipe alias tabBarang dengan maksimum elemen NMAX
type tabBarang [NMAX]tBarang

func main() {
	// Deklarasi variabel jualan bertipe tabBarang
	var jualan tabBarang

	// Deklrasi variabel banyak barang bertipe integer
	var nBarang int

	// Input jumlah barang
	fmt.Scan(&nBarang)

	// Memanggil prosedur bacaBarang
	bacaBarang(&jualan, &nBarang)

	// Memanggil prosedur cetakBarang
	cetakBarang(jualan, nBarang)

	// Memanggil fungsi totalAset
	fmt.Println("Total Aset:", totalAset(jualan, nBarang))

	// Memanggil fungsi mencari harga barang termahal
	fmt.Println("Nama Barang Termahal:", namaHargaBarangTermahal(jualan, nBarang))
}

func bacaBarang(tB *tabBarang, n *int) {
	/*
		IS: Array tB sebanyak n elemen terdefinisi sembarang
		Proses : Membaca n, dan membaca n barang dengan field: id, nama, harga, stok.
		         Jumlah n tidak boleh melebihi NMAX
		FS: ArraytB sebanyak n elemen terdefinisi
	*/
	var inputId, inputNama string
	var inputHarga, inputStok int
	// var nAsli digunakan untuk jumlah iterasi
	var nAsli int = *n

	// ubah n jika lebih besar dari NMAX
	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < nAsli; i++ {
		fmt.Scan(&inputId)
		fmt.Scan(&inputNama)
		fmt.Scan(&inputHarga)
		fmt.Scan(&inputStok)

		// masukkan input ke array hanya jika jumlah iterasi masih di bawah n atau NMAX
		if i < *n {
			tB[i].id = inputId
			tB[i].nama = inputNama
			tB[i].harga = inputHarga
			tB[i].stok = inputStok
		}
	}
}

func cetakBarang(tB tabBarang, n int) {
	/*
		IS: Array tB terdefinisi sebanyak n elemen
		FS: Tercetak data barang dengan format:
			"Data Barang:
			 <id> <nama> <harga> <stok>
			 ...   ...	  ....	  ... "
	*/
	fmt.Println("Data Barang:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", tB[i].id, tB[i].nama, tB[i].harga, tB[i].stok)
	}
}

func totalAset(tB tabBarang, n int) int {
	/* Diberikan array tB sebanyak n elemen, fungsi akan mengembalikan total aset yang
	   merupakan total perkalian seluruh harga barang dengan stoknya */
	var totalAset int = 0

	for i := 0; i < n; i++ {
		totalAset += tB[i].harga * tB[i].stok
	}

	return totalAset
}

func namaHargaBarangTermahal(tB tabBarang, n int) string {
	/* Diberikan array tB sebanyak n elemen, fungsi mengembalikan nama barang dengan harga
	   termahal */
	// var barangTermahal untuk menampung tBarang dengan harga termahal
	// inisialisasi nilai awal dengan tB index ke-0
	var barangTermahal tBarang = tB[0]

	for i := 0; i < n; i++ {
		if tB[i].harga > barangTermahal.harga {
			barangTermahal = tB[i]
		}
	}

	// hanya kembalikan nilai nama-Nya
	return barangTermahal.nama
}
