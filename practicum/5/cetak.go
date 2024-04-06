package main

import "fmt"

/*  Buatlah konversi pola end-tail recursion untuk prosedur mencetak N bilangan
    asli pertama yang merupakan kelipatan 2 atau 3
*/

func main() {
	var suku int
	fmt.Scan(&suku)
	// 	cetak_iterate(suku)
	// 	fmt.Println()
	cetak_recurse_master(suku)
	fmt.Println()
}

func cetak_iterate(N int) {
	var i int
	i = 1
	for i <= N {
		if i%2 == 0 || i%3 == 0 {
			fmt.Print(i, " ")
		}
		i++
	}
}

func cetak_recurse_master(N int) {
	var i int = 1
	cetak_recurse(N, i)
}

func cetak_recurse(N, i int) {
	if !(i < N) {
		if i%2 == 0 || i%3 == 0 {
			fmt.Print(i, " ")
		}
	} else {
		if i%2 == 0 || i%3 == 0 {
			fmt.Print(i, " ")
		}
		i++
		cetak_recurse(N, i)
	}
}
