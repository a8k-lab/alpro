package main

import "fmt"

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm++
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd++
	}
}

func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk++
	}
}

func hitungJumGolKegolanSelisih(g, k int, jg, jk, jsg *int) {
	*jg += g
	*jk += k
	*jsg += g - k
}

func hitungJumPoint(jp *int) {

}

func main() {
	var N, gol, kemasukan int
	var jumGol, jumKemasukan, jumSelisihGol, jumMenang, jumDraw, jumKalah int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&gol, &kemasukan)

		hitungMenang(gol, kemasukan, &jumMenang)
		hitungDraw(gol, kemasukan, &jumDraw)
		hitungKalah(gol, kemasukan, &jumKalah)
		hitungJumGolKegolanSelisih(gol, kemasukan, &jumGol, &jumKemasukan, &jumSelisihGol)
	}
}
