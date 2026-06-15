package main

import "fmt"

func city() string {
	return "HYDERABAD"
}
func age(age int) string {
	if age >= 18 {
		return "you can vote"
	} else {
		return "you can't"
	}
}
func vote() float64 {
	return (64.0)

}

func main() {
	fmt.Println(city())
	fmt.Println(age(26))
	fmt.Println(vote())

}
