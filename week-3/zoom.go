package main

import "fmt"

func main() {
	var width, height, scale int
	var operation string

	fmt.Scan(&width, &height)
	fmt.Scan(&operation, &scale)

	if operation == "+" {
		zoomIn(&width, &height, scale)
	} else if operation == "-" {
		zoomOut(&width, &height, scale)
	} else {
		return
	}

	fmt.Println(width, height)
}

func zoomIn(width, height *int, scale int) {
	*width = *width * scale
	*height = *height * scale
}

func zoomOut(width, height *int, scale int) {
	*width = *width / scale
	*height = *height / scale
}
