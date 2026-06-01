package main

import "fmt"

func main() {

	//creating a slice

	colors := []string{"red", "green", "blue", "black"}

	// add elements to slices
	colors = append(colors, "white")
	//slicing an existing slice/array[start:end]4 is execlusive

	Primarycolors := colors[0:5]

	// initializing empty slice make([] type,length,capacity)

	numbers := make([]int, 4, 7)

	fmt.Println("colors:", colors)
	fmt.Println("Primary colors:", Primarycolors)
	fmt.Println("Numbers length & Capacity", len(numbers), cap(numbers))

}
