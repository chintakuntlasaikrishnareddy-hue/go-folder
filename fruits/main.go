package main

import "fmt"

func main() {
	//array of a 3 strings

	var fruits [4]string
	fruits[0] = "mango"
	fruits[1] = "apple"
	fruits[2] = "strawberry"
	fruits[3] = "grapes"

	// define an array in single line

	primes := [5]int{2, 3, 5, 7, 11}

	fmt.Println("fruits:", fruits)
	fmt.Println("primes:", primes)
	fmt.Println("length of primes:", len(primes))

}
