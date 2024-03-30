package main

import "fmt"

func main() {
	var celciusInit, celciusEnd, step float64

	fmt.Scan(&celciusInit, &celciusEnd, &step)

	fmt.Println("Celcius", "Reamur", "Fahrenheit", "Kelvin")

	for celciusInit <= celciusEnd {
		C := celciusInit
		R := reamur(C)
		F := fahrenheit(C)
		K := kelvin(C)

		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", C, R, F, K)
		celciusInit += step
	}
}

func reamur(celcius float64) float64 {
	return celcius * (4.0 / 5.0)
}

func fahrenheit(celcius float64) float64 {
	return celcius*(9.0/5.0) + 32
}

func kelvin(celcius float64) float64 {
	return celcius + 273.00
}
