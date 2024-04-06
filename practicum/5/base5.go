package main

// basis 5 to desimal

import "fmt"

func main() {
	var bil, des int
	fmt.Scan(&bil)
	des = basis5todesimal(bil, 0)
	fmt.Println(des)
}

func basis5todesimal(number, posisi int) int {
	var lastDigit, desimal int
	if number == 0 {
		return 0
	} else {
		lastDigit = number % 10
		desimal = lastDigit*pangkat(5, posisi) + basis5todesimal(number/10, posisi+1)
		return desimal
	}
}

func pangkat(basis, eksponen int) int {
	var i, hasil int
	hasil = 1
	for i = 0; i < eksponen; i++ {
		hasil = hasil * basis
	}
	return hasil
}
