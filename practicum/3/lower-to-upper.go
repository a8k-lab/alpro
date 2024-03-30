package main

import "fmt"

func main() {
	var loweChar, upperChar byte

	fmt.Scanf("%c", &loweChar)

	upperChar = lowToUpper(loweChar)
	fmt.Println(string(upperChar))
}

func lowToUpper(kar byte) byte {
	if kar >= 'a' && kar <= 'z' {
		return kar - 32
	}

	return kar
}
