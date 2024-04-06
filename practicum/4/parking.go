package main

import "fmt"

// variabel global
var jenisKendaraan string
var jam1, menit1, detik1 int
var jam2, menit2, detik2 int
var totalUang int

func main() {
	var pilih int
	var ulang bool = true

	for ulang {
		menu()
		fmt.Print("Pilih (1/2/3/4/5)? ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			inputKendaraanMasuk()
		} else if pilih == 2 {
			inputKendaraanKeluar()
		} else if pilih == 3 {
			hitungBiayaParkir()
		} else if pilih == 4 {
			cetakTotalUang()
		} else {
			ulang = false
		}
	}
}

func menu() {
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Input Kendaraan Masuk")
	fmt.Println("2. Input Kendaraan keluar")
	fmt.Println("3. Hitung Biaya Parkir")
	fmt.Println("4. Cetak Total Uang")
	fmt.Println("5. Exit")
	fmt.Println("-----------------------")
}

func inputKendaraanMasuk() {
	fmt.Print("Masukkan jenis kendaraan (mobil/motor): ")
	fmt.Scan(&jenisKendaraan)
	fmt.Print("Masukkan jam, menit, detik kendaraan masuk: ")
	fmt.Scan(&jam1, &menit1, &detik1)
}

func inputKendaraanKeluar() {
	fmt.Print("Masukkan jam, menit, detik kendaraan keluar: ")
	fmt.Scan(&jam2, &menit2, &detik2)
}

func hitungBiayaParkir() {
	totalJam := jam2 - jam1
	totalMenit := menit2 - menit1
	totalDetik := detik2 - detik1

	totalMenit += totalDetik / 60 // setiap 60 dari detik dijadikan 1 dan ditambah ke menit
	totalJam += totalMenit / 60   // setiap 60 dari menit dijadikan 1 dan ditambah ke jam

	if jenisKendaraan == "mobil" {
		totalUang += 5000 // 1 jam pertama
		if totalJam > 1 {
			totalUang += 3000 * (totalJam - 1)
		}
	} else {
		totalUang += 2000 // 1 jam pertama
		if totalJam > 1 {
			totalUang += 1000 * (totalJam - 1)
		}
	}

	reset()
}

func reset() {
	jenisKendaraan = ""
	jam1 = 0
	menit1 = 0
	detik1 = 0
	jam2 = 0
	menit2 = 0
	detik2 = 0
}

func cetakTotalUang() {
	fmt.Printf("Total uang: Rp. %d,-\n", totalUang)
}
