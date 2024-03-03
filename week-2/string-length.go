package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(stringLength(str))
}

func stringLength(str string) int {
	return len(str)
}
