package main

import "fmt"

func main() {
	var jam, menit int
	var member bool
	var total float64

	fmt.Scan(&jam, &menit, &member)
	hitungSewa(jam, menit, member, &total)
	fmt.Println(total)
}

func durasi(jam, menit int) int {
	/* mengembalikan durasi dalam jam, apabila diketahui lama sewa dalam jam dan menit.
	   Sewa di atas 1 jam, maka kelebihan menit kurang dari 10 menit tidak dihitung penambahan jam*/
	var hasil int = 0
	if jam >= 1 {
		hasil += jam
	}
	if menit >= 10 || (jam < 1 && menit < 10 && menit > 1) {
		hasil += 1
	}

	return hasil
}

func potongan(durasi int, tarif int) float64 {
	/* mengembalikan besarnya potongan apabila diketahui durasi dalam jam dan tarif perjamnya*/
	if durasi > 3 {
		return float64(tarif) * 0.10
	}
	return 0
}

func tarif(member bool) int {
	/* mengembalikan tarif sewa sesuai dengan status membernya */
	if member {
		return 3500
	}
	return 5000
}

func hitungSewa(jam, menit int, member bool, biaya *float64) {
	/* I.S. terdefinisi durasi sewa dalam jam dan menit, serta status membershipnya
	   F.S. biaya berisi total biaya sewa setelah dipotong diskon*/
	var totalDurasi, totalTarif int
	var totalPotongan float64

	totalDurasi = durasi(jam, menit)
	totalTarif = tarif(member) * totalDurasi
	totalPotongan = potongan(totalDurasi, totalTarif)

	*biaya = float64(totalTarif) - totalPotongan
	fmt.Println("Durasi:", totalDurasi)
}
