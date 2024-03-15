package main

import "fmt"

func main() {
	var clo1, clo2, clo3 float64

	fmt.Scan(&clo1, &clo2, &clo3)
	for clo1 >= 0 && clo2 >= 0 && clo3 >= 0 {
		cekCLO(&clo1, &clo2, &clo3)
	}
}

func cekCLO(clo1, clo2, clo3 *float64) {
	if *clo1 > 50.00 && *clo2 > 50.00 && *clo3 > 50.00 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	fmt.Scan(clo1, clo2, clo3)
}
