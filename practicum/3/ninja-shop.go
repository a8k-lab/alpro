package main

import "fmt"

func ganjil(a, b, c, d, e int) bool {
	return a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0
}

func genap(a, b, c, d, e int) bool {
	return a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0
}

func diskon(member string, c, d, e int) float64 {
	totalPoin := float64(c + d + e)
	nilaiDiskon := 1.7 * totalPoin

	if member == "yes" {
		nilaiDiskon += 0.15 * nilaiDiskon
	}
	if totalPoin/nilaiDiskon >= 2 {
		nilaiDiskon = 0.5 * totalPoin
	}

	return nilaiDiskon
}

func cashback(member string, a, b, c int) float64 {
	totalPoin := float64(a + b + c)
	nilaiCashback := 3.1 * totalPoin

	if member == "yes" {
		nilaiCashback += 0.15 * nilaiCashback
	}
	if nilaiCashback >= 35 {
		nilaiCashback = 35.00
	}

	return nilaiCashback
}

func main() {
	var member string
	var a, b, c, d, e int
	var totalDiskon, totalCashback float64 = 0, 0

	fmt.Scan(&member)
	fmt.Scan(&a, &b, &c, &d, &e)

	if ganjil(a, b, c, d, e) {
		totalDiskon = diskon(member, c, d, e)
	} else if genap(a, b, c, d, e) {
		totalCashback = cashback(member, a, b, c)
	} else {
		totalDiskon = diskon(member, c, d, e)
		totalCashback = cashback(member, a, b, c)
	}

	fmt.Printf("%.2f %.2f\n", totalDiskon, totalCashback)
}
