package main

import "fmt"

func main() {

	// SWITCH

	/* status := "active"
	switch status {
	case "active":
		fmt.Println("service is running")

	case "down":
		fmt.Println("Alert!Service is Unreachable")

	case "maintain":
		fmt.Println("Service is under maintaince")

	default:
		fmt.Println("UNKNOWN STATUS")

	}

	day := "tuesday"

	switch day {
	case "monday", "tuesday", "wednesday", "thuresday", "friday":
		fmt.Println("Its a Weekday")

	case "saturday", "sunday":
		fmt.Println("Its a wekend")

	default:
		fmt.Println("Invalid Day")
	}*/

	//fallThrough

	step := 1

	switch step {
	case 1:
		fmt.Println("step 1: fetching metrics")

		fallthrough

	case 2:
		fmt.Println("step 2: parsing JSON")
		fallthrough

	case 3:
		fmt.Println("step3: preparing to DataBase")

	}
}
