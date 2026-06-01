package main

import (
	"fmt"
)

func main() {
	//Area calculator
	// given area of rectangle
	// A=L*W

	Length := 25.0
	width := 8.0

	Area := Length * width

	fmt.Printf("Length: %.3f\n", Length)
	fmt.Printf("Width: %.3f\n", width)
	fmt.Printf("Area: %.3f\n", Area)
}
