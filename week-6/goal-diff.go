package main

import "fmt"

func main() {
	type KlubBola struct {
		nama              string
		point, selisihGol int
	}
	var inputKlub KlubBola
	var totalSelisihGol int
	var avgSelisihGol float64

	for i := 0; i < 3; i++ {
		fmt.Scan(&inputKlub.nama, &inputKlub.point, &inputKlub.selisihGol)
		totalSelisihGol += inputKlub.selisihGol
	}
	avgSelisihGol = float64(totalSelisihGol) / 3
	fmt.Println(avgSelisihGol)
}
