package main

import "fmt"

type employee struct {
	Id       int
	Name     string
	position string
	Active   bool
}

func main() {

	emp1 := employee{
		Id:       105,
		Name:     "Sai",
		position: "HR",
		Active:   true,
	}

	fmt.Println("employee Name:", emp1.Name)
	fmt.Println("employee position:", emp1.position)
	fmt.Println("employee Id:", emp1.Id)

	emp1.Active = true

	// Ananymous Struct

	point := struct {
		x int
		y int
	}{x: 10, y: 25}

	fmt.Println("point:", point)

}
