package main

import "fmt"

func main() {

	// map [keytype] valuetype
	// create a map of inventory
	inventory := make(map[string]int)

	inventory["mango"] = 52
	inventory["grapes"] = 22

	// short-hand
	UserRoles := map[string]string{
		"Sai":     "Admin",
		"krishna": "User",
	}
	fmt.Println("inventory:", inventory)
	fmt.Println("UserRoles:", UserRoles)

}
