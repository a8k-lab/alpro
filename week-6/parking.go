package main

import "fmt"

type Kendaraan struct {
	jenis string
	tarif int
}

var kendaraan Kendaraan

func hitungTarif(totalTarif *int) {
	var durasi int
	for {
		fmt.Scan(&kendaraan.jenis, &durasi)

		if kendaraan.jenis == "motor" {
			kendaraan.tarif = 1000
		} else if kendaraan.jenis == "mobil" {
			kendaraan.tarif = 5000
		} else {
			break
		}
		if durasi < 1 {
			break
		}

		*totalTarif += durasi * kendaraan.tarif
	}
}

func main() {
	var totalTarif int

	hitungTarif(&totalTarif)
	fmt.Println(totalTarif)
}
