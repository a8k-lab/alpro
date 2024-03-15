package main

import "fmt"

func main() {
	var moneyInput, k10, k2, k1 int

	fmt.Scan(&moneyInput)
	denom(moneyInput, &k10, &k2, &k1)

	fmt.Println(k10, "lembar 10000")
	fmt.Println(k2, "lembar 2000")
	fmt.Println(k1, "lembar 1000")
}

func denom(money int, k10, k2, k1 *int) {
	if money >= 10_000 {
		*k10 = money / 10_000
		money = money % 10_000
	} else {
		*k10 = 0
	}

	if money >= 2000 {
		*k2 = money / 2000
		money = money % 2000
	} else {
		*k2 = 0
	}

	if money >= 1000 {
		*k1 = money / 1000
		money = money % 1000
	} else {
		*k1 = 0
	}
}
