package main

import "fmt"

func main() {
	var celcius, reamur, fahrenheit, kelvin float64

	fmt.Scan(&celcius)
	tempConverter(celcius, &reamur, &fahrenheit, &kelvin)

	fmt.Print(reamur, "R ")
	fmt.Print(fahrenheit, "F ")
	fmt.Print(kelvin, "K")
}

func tempConverter(celcius float64, reamur, fahrenheit, kelvin *float64) {
	*reamur = celcius * (4.0 / 5.0)
	*fahrenheit = celcius*(9.0/5.0) + 32
	*kelvin = celcius + 273.15
}
