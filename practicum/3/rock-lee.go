package main

import "fmt"

func pukulan(P, X, Y, Z, H int) int {
	totalPukulan := 0
	for i := 1; i <= H; i++ {
		if (i%X == 0 || i%Y == 0) && i%Z != 0 {
			totalPukulan += P
		}
	}

	return totalPukulan
}

func tendangan(T, X, Y, Z, H int) int {
	totalTendangan := 0
	for i := 1; i <= H; i++ {
		if (i%X == 0 && i%Y == 1) && i%Z != 0 {
			totalTendangan += T
		}
	}

	return totalTendangan
}

func delapan_gerbang(G, X, Y, Z, H int) int {
	totalJam := 0
	for i := 1; i <= H; i++ {
		if (i%X == 0 && i%Y != 0) && i%Z != 0 {
			totalJam += G
		}
	}

	return totalJam
}

func main() {
	var P, T, G, X, Y, Z, H int
	var totalPukulan, totalTendangan, totalJamDelapanGerbang int
	fmt.Scan(&P, &T, &G, &X, &Y, &Z, &H)

	totalPukulan = pukulan(P, X, Y, Z, H)
	totalTendangan = tendangan(T, X, Y, Z, H)
	totalJamDelapanGerbang = delapan_gerbang(G, X, Y, Z, H)

	fmt.Println(totalPukulan, totalTendangan, totalJamDelapanGerbang)
}
