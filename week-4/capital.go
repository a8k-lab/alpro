package main

import "fmt"

func main() {
	var text string
	var counter, index int
	index = 0

	fmt.Scanln(&text)
	counter = capitalCounter(text, index)
	fmt.Print(counter)
}

func capitalCounter(text string, index int) int {
	if index == len(text) {
		return 0
	}

	if 'A' <= text[index] && text[index] <= 'Z' {
		return 1 + capitalCounter(text, index)
	}

	return capitalCounter(text, index)
}
