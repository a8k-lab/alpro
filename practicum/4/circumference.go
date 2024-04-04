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
	var resultArr = [20][8]float64{}
	var resultLen int

	fmt.Scan(&R, &S)
	if R == 0 && S == 0 {
		return
	}

	for i := 0; R != 0 && S != 0; i++ {
		hitungLuasKelilingLingkaran(R, &LL, &KL)
		hitungLuasKelilingPersegi(S, &LP, &KP)
		hitungTotal(LL, LP, KL, KP, &TL, &TP)

		resultArr[i] = [8]float64{R, S, LL, LP, KL, KP, TL, TP}
		resultLen = i + 1

		fmt.Scan(&R, &S)
	}

	fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s \n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")

	for _, v := range resultArr[:resultLen] {
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f \n", v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7])
	}
}
