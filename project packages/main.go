package main

import (
	"fmt"
	"projectpackages/bank"
)

func main() {
	account := bank.NewAccount("sai", 500.00)
	fmt.Println("Account Number:", account.Owner)
	account.Deposit(1500.00)
	fmt.Printf("current Balance:%.2f\n", account.GetBalance())

}
