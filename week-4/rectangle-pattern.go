package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)
	printPolaBilangan(1, 1, N)
}

func printPolaBilangan(baris, kolom, n int) {
	if n >= baris {
		if n >= kolom {
			if baris == 1 || baris == n || kolom == 1 || kolom == n { // Pengecekan Pengulangan kolom dan baris dengan operasi perbandingan,
				// hint : terdapat beberapa operasi perbandingan
				fmt.Print("*", " ")
			} else {
				fmt.Print("  ")
			}
			printPolaBilangan(baris, kolom+1, n) // gunakan fungsi rekursif pada baris ini
		} else {
			fmt.Println()
			printPolaBilangan(baris+1, 1, n) // gunakan fungsi rekursif pada baris ini
		}
	}
}
