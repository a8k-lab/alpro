package main

import "fmt"

// Tipe bentukan titik dengan field x dan y
type titik struct {
	x, y float64
}

// tipe bentukan struct lingkaran dengan field titik_pusat bertipe titik dan radius, luas, keliling bertip float64
type lingkaran struct {
	titik_pusat            titik
	radius, luas, keliling float64
}

func main() {
	var c1, c2 lingkaran
	var x, y, radius float64

	fmt.Scan(&x, &y, &radius)
	// membuat lingkaran baru c1
	c1 = lingkaran_baru(x, y, radius)

	fmt.Scan(&x, &y, &radius)
	// membuat lingkaran baru c2
	c2 = lingkaran_baru(x, y, radius)

	// Menghitung luas c1
	hitung_luas(&c1)

	// Menghitung luas c2
	hitung_luas(&c2)

	// Menghitung keliling c1
	hitung_keliling(&c1)

	// Menghitung keliling c2
	hitung_keliling(&c2)

	// cetak data
	cetak_data(c1)
	cetak_data(c2)
}

func lingkaran_baru(x, y, r float64) lingkaran {
	/* Mengembalikan lingkaran dengan titik pusat x, y dan radius r */
	var l lingkaran
	var tp titik

	tp.x = x
	tp.y = y
	l.titik_pusat = tp
	l.radius = r

	return l
}

func hitung_luas(l *lingkaran) {
	/* Mengembalikan luas lingkaran (l.luas) dengan pi = 3.14 */
	const pi float64 = 3.14
	l.luas = pi * l.radius * l.radius
}

func hitung_keliling(l *lingkaran) {
	/* Mengembalikan keliling lingkaran (l.keliling) dengan pi = 3.14 */
	const pi float64 = 3.14
	l.keliling = 2 * pi * l.radius
}

func cetak_data(l lingkaran) {
	/* 	IS: Lingkaran l terdefinisi
		FS: Tercetak data lingkaran l dengan format:
	        "Lingkaran berpusat di titik (<x, y>) dan berradius <radius> memiliki luas <luas> dan keliling
	         sebesar <keliling>"
		    Luas dan keliling dibulatkan 1 digit di belakang koma
	*/
	fmt.Printf("Lingkaran berpusat di titik (%g, %g) dan berradius %g ", l.titik_pusat.x, l.titik_pusat.y, l.radius)
	fmt.Printf("memiliki luas %.1f dan keliling sebesar %.1f\n", l.luas, l.keliling)
}
