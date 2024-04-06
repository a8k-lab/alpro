package main

import (
	"fmt"
	"math"
)

func konversiDerajatkeRadian(T float64) float64 {
	return (T * math.Pi) / 180
}

func waktu(V, T float64) float64 {
	return V * math.Sin(konversiDerajatkeRadian(T)) / 9.8
}

func jarak(V, T float64) float64 {
	return math.Pow(V, 2) * math.Sin(2*konversiDerajatkeRadian(T)) / 9.8
}

func ketinggian(V, T float64) float64 {
	return math.Pow(V, 2) * math.Pow(math.Sin(konversiDerajatkeRadian(T)), 2) / (2 * 9.8)
}

func main() {
	var V, T float64

	fmt.Scan(&V, &T)
	fmt.Printf("%.2f %.2f %.2f\n", waktu(V, T), jarak(V, T), ketinggian(V, T))
}
