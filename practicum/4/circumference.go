package main

import "fmt"

func hitungLuasKelilingLingkaran(radius float64, luas, keliling *float64) {
	*luas = 3.14 * radius * radius
	*keliling = 2 * 3.14 * radius
}

func hitungLuasKelilingPersegi(sisi float64, luas, keliling *float64) {
	*luas = sisi * sisi
	*keliling = 4 * sisi
}

func hitungTotal(lLingkaran, lPersegi, kLingkaran, kPersegi float64, tLuas, tKeliling *float64) {
	*tLuas = lLingkaran + lPersegi
	*tKeliling = kLingkaran + kPersegi
}

func main() {
	var R, S, LL, LP, KL, KP, TL, TP float64
	R = -1
	S = -1

	fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s \n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	for R != 0 && S != 0 {
		fmt.Scan(&R, &S)

		hitungLuasKelilingLingkaran(R, &LL, &KL)
		hitungLuasKelilingPersegi(S, &LP, &KP)
		hitungTotal(LL, LP, KL, KP, &TL, &TP)

		fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s \n", R, S, LL, LP, KL, KP, TL, TP)
	}
}
