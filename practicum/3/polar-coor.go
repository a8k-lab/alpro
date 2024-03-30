package main

import (
	"fmt"
	"math"
)

func main() {
	var R, T, X, Y float64

	fmt.Scan(&R, &T)
	X = getXLength(R, T)
	Y = getYLength(R, T)
	fmt.Printf("%.2f %.2f \n", X, Y)
}

func getXLength(R, T float64) float64 {
	T = T * math.Pi / 180
	return R * math.Cos(T)
}

func getYLength(R, T float64) float64 {
	T = T * math.Pi / 180
	return R * math.Sin(T)
}
