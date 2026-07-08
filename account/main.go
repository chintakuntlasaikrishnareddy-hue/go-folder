package main

import "fmt"

type Account struct {
	Name    string
	Balance int32
	Status  bool
}

func main() {

	accounts := []Account{
		{
			Name:    "Sai",
			Balance: 17101999,
			Status:  true,
		},
		{
			Name:    "Harsha",
			Balance: 25082002,
			Status:  false,
		},
		{
			Name:    "Ammu",
			Balance: 3122007,
			Status:  true,
		},
	}

	searchName := "Harsha"
	found := false

	for _, acc := range accounts {
		if acc.Name == searchName {
			fmt.Println("--- Searched Account Details ---")
			fmt.Println("Account Holder Name:", acc.Name)
			fmt.Println("Balance:", acc.Balance)

			if acc.Status == true {
				fmt.Println("Active Status: Active")
			} else {
				fmt.Println("Active Status: Not Active")
			}

			found = true
			break
		}
	}

	if !found {
		fmt.Println("Account Not Found")
	}

	fmt.Println("\n--- Other Accounts Details ---")
	for _, acc := range accounts {
		if acc.Name != searchName {
			fmt.Println("Account Holder Name:", acc.Name)
			fmt.Println("Balance:", acc.Balance)

			if acc.Status == true {
				fmt.Println("Active Status: Active")
			} else {
				fmt.Println("Active Status: Not Active")
			}
			fmt.Println("---------------------------")
		}
	}
}
