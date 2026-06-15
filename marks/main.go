package main

import "fmt"

func percentage(obtainedmarks float64, totalmarks float64) float64 {
	return (obtainedmarks / totalmarks) * 100

}
func grade(grade float64) string {
	if grade >= 85 {
		return "FIRST CLASS"
	} else if grade >= 65 {
		return "SECOND CLASS"
	} else if grade >= 45 {
		return "THIRD CLASS"
	} else {
		return "FAIL"
	}
}
func main() {
	result := percentage(520, 600)
	fmt.Println(result)
	fmt.Println(grade(result))
}
