package main

import "fmt"

func jumlahBilangan(bilangan *int) {
	/**I.S Terdefinisi sebuah variabel integer jumlah bernilai nol
	F.S Prosedur akan menjumlahkan bilangan yang diinputkan secara terus-menerus
	hingga terbaca bilangan negatif dan ditampung oleh variabel jumlah
	*/
	var angkaBaru int = 0

	for angkaBaru >= 0 {
		*bilangan += angkaBaru
		fmt.Scan(&angkaBaru)
	}
}

func main() {
	var jumlah int = 0

	jumlahBilangan(&jumlah)
	fmt.Println(jumlah)
}
