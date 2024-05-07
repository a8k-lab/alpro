package main

import "fmt"

type pegawai struct {
	nama                    string
	masa_kerja              int
	gaji_pokok, besar_bonus float64
}

func hitung_bonus(p *pegawai) {
	if p.masa_kerja >= 10 {
		p.besar_bonus = p.gaji_pokok * 1.5
	} else if p.masa_kerja >= 5 {
		p.besar_bonus = p.gaji_pokok * 0.75
	} else {
		p.besar_bonus = p.gaji_pokok * 0.5
	}
}

func main() {
	var pegawai1 pegawai

	fmt.Scan(&pegawai1.nama, &pegawai1.gaji_pokok, &pegawai1.masa_kerja)

	hitung_bonus(&pegawai1)

	fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp %.0f\n", pegawai1.nama, pegawai1.besar_bonus)
}
