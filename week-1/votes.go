package main

import "fmt"

func main() {
	var lolos bool
	var totalYes int
	var vote1, vote2, vote3 string

	totalYes = 0
	fmt.Scan(&vote1, &vote2, &vote3)

	if vote1 == "yes" {
		totalYes++
	}
	if vote2 == "yes" {
		totalYes++
	}
	if vote3 == "yes" {
		totalYes++
	}

	lolos = totalYes >= 2
	if lolos {
		fmt.Println("lolos")
	} else {
		fmt.Println("tidak lolos")
	}
}
