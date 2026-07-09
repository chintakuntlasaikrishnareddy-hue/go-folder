package main

import "fmt"

type MetroStation struct {
	StationName       string
	FloatingCustomers int
	Active            bool
}

func main() {

	stations := []MetroStation{
		{
			StationName:       "MARATHALLI",
			FloatingCustomers: 25000,
			Active:            true,
		},
		{
			StationName:       "KUNDHANHALLI",
			FloatingCustomers: 32456,
			Active:            true,
		},
		{
			StationName:       "WHITEFIELD",
			FloatingCustomers: 2500,
			Active:            false,
		},
		{
			StationName:       "BENGENAHALLI",
			FloatingCustomers: 52458,
			Active:            false,
		},
	}

	searchName := "KUNDHANHALLI"
	found := false

	for _, station := range stations {

		if station.StationName == searchName {

			fmt.Println("STATION FOUND")
			fmt.Println("Station Name:", station.StationName)
			fmt.Println("Floating Customers:", station.FloatingCustomers)

			if station.Active {
				fmt.Println("Status: Active")
			} else {
				fmt.Println("Status: Inactive")
			}

			found = true
			break
		}
	}

	if !found {
		fmt.Println("Station Not Found")
	}

	fmt.Println("\n----- OTHER STATIONS -----")

	for _, station := range stations {

		if station.StationName == searchName {
			continue
		}

		fmt.Println("Station Name:", station.StationName)
		fmt.Println("Floating Customers:", station.FloatingCustomers)

		if station.Active {
			fmt.Println("Status: Active")
		} else {
			fmt.Println("Status: Inactive")
		}

		fmt.Println("----------------------")
	}
}
