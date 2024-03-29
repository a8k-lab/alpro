package main

import "fmt"

func main() {
	var N int
	var result string
	fmt.Scan(&N)
	result = printStars(N, 1, 0, "")
	fmt.Print(result)
}

func printStars(totalStars, current, zero int, resultStars string) string {
	if totalStars%2 == 0 {
		return printStars(totalStars-1, current, zero, resultStars)
	}

	if current%2 != 0 {
		for i := 0; i < current; i++ {
			resultStars += "*"
		}
		resultStars += "\n"
	}

	if current < totalStars {
		return printStars(totalStars, current+1, 0, resultStars)
	}

	return resultStars
}
