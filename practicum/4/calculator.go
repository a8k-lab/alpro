package main

import "fmt"

func menu() {
	fmt.Println("-----------------------")
	fmt.Printf("%15s\n", "M E N U")
	fmt.Println("-----------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")
}

func hitungJumlah() {
	var a, b int

	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan: ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil penjumlahan:", a+b)
}

func hitungKali() {
	var a, b int

	fmt.Print("Masukkan dua bilangan yang akan dikalikan: ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil perkalian:", a*b)
}

func hitungBagi() {
	var a, b float64

	fmt.Print("Masukkan dua bilangan yang akan dibagikan: ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil pembagian:", a/b)
}

func main() {
	var pilih int

	for pilih != 4 {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			hitungJumlah()
		} else if pilih == 2 {
			hitungKali()
		} else if pilih == 3 {
			hitungBagi()
		}
	}
}
