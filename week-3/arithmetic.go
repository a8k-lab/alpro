package main

import "fmt"

func main() {
	var u1, u2, u3 int

	fmt.Scan(&u1, &u2, &u3)
	arithmeticRow(u1, u2, u3)
}

func arithmeticRow(u1, u2, u3 int) {
	if u2-u1 == u3-u2 {
		fmt.Println("y")
	} else {
		fmt.Println("t")
	}
}
