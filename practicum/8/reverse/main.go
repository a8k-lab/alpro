package main

import "fmt"

const NMAX int = 20

type tabChar [NMAX]rune

func main() {
	var kata tabChar
	var nChar int

	fmt.Scan(&nChar)
	baca(&kata, &nChar)
	cetak(kata, nChar)
}

func baca(k *tabChar, n *int) {
	var inputKata string
	fmt.Scan(&inputKata)

	// jika n nilainya lebih dari NMAX
	// maka n diubah nilainya menjadi NMAX (10)
	if *n > NMAX {
		*n = NMAX
	}

	// melakukan perulangan ke setiap karakter dari hasil input
	for i := 0; i < *n; i++ {
		k[i] = rune(inputKata[i])
	}
}

func cetak(k tabChar, n int) {
	for i := n - 1; i >= 0; i-- {
		// menggunakan format %c untuk mengubah karakter ke ASCII
		fmt.Printf("%c", k[i])
	}
	fmt.Println()
}
