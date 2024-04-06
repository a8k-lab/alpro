package main

import "fmt"

func main() {
	var M, N int

	fmt.Printf("Berapa kekuatan awal gerbang? ")
	fmt.Scan(&M)
	fmt.Printf("Berapa daya rusak balok? ")
	fmt.Scan(&N)

	fmt.Println("Kekuatan awal gerbang", M)
	dobrakPintu(M, N, 0)
}

func dobrakPintu(M, N, O int) {
	O += 1
	if M > N {
		M -= N
		fmt.Println("dobrak, sisa kekuatan", M)
		dobrakPintu(M, N, O)
	} else {
		fmt.Println("dobrak, gerbang jebol")
		fmt.Printf("Pasukan berusaha sebanyak %d kali\n", O)
	}
}
