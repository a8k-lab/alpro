package main

import "fmt"

func vowel(kar byte) bool {
	return kar == 'a' || kar == 'i' || kar == 'u' || kar == 'e' || kar == 'o' || kar == 'A' || kar == 'I' || kar == 'U' || kar == 'E' || kar == 'O'
}

func main() {
	var kar byte
	var totalVocal int = 0

	fmt.Scanf("%c", &kar)
	for kar != '.' {
		if vowel(kar) {
			totalVocal++
		}

		fmt.Scanf("%c", &kar)
	}

	fmt.Println("Jumlah vokal adalah", totalVocal)
}
