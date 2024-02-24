package main

import "fmt"

func main() {
	var a, b, c, d int
	var x, y int

	fmt.Scan(&a, &b, &c, &d)

	// operasi untuk penjumlahan
	x = (a * d) + (b * c)
	y = b * d
	fmt.Println("penjumlahan:", x, "/", y)

	// operasi untuk pengurangan
	x = (a * d) - (b * c)
	y = b * d
	fmt.Println("pengurangan:", x, "/", y)

	// operasi untuk perkalian
	x = a * c
	y = b * d
	fmt.Println("perkalian:", x, "/", y)

	// operasi untuk pembagian
	x = a * d
	y = b * c
	fmt.Println("pembagian:", x, "/", y)
}
