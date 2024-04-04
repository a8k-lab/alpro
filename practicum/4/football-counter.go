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

func hitungJumPoint(jm, jd int, jp *int) {
	*jp = jm*3 + jd
}

func main() {
	var N, gol, kegolan int
	var jumGol, jumKegolan, jumSelisihGol, jumMenang, jumDraw, jumKalah, jumPoin int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&gol, &kegolan)

		hitungMenang(gol, kegolan, &jumMenang)
		hitungDraw(gol, kegolan, &jumDraw)
		hitungKalah(gol, kegolan, &jumKalah)
		hitungJumGolKegolanSelisih(gol, kegolan, &jumGol, &jumKegolan, &jumSelisihGol)
		hitungJumPoint(jumMenang, jumDraw, &jumPoin)
	}

	jumMain := jumMenang + jumDraw + jumKalah
	fmt.Println(jumMain, jumMenang, jumDraw, jumKalah, jumGol, jumKegolan, jumSelisihGol, jumPoin)
}
