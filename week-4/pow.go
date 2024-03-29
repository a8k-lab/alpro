package main

import "fmt"

func main() {
	var x, y, xPowY int
	fmt.Scan(&x, &y)

	xPowY = pow(x, y)

	fmt.Println(xPowY)
}

func pow(x, y int) int {
	hasil := x

	if x < 1 {
		return 0
	} else if y > 1 {
		hasil *= x
		pow(x, y-1)
	}

	return hasil
}
