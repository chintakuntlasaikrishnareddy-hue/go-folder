package main

import "fmt"

type tickets struct {
	name       string
	passportno string
	price      float64
	valid      bool
}

func main() {
	ticket1 := tickets{
		name:       "sai",
		passportno: "P58246",
		price:      5645.00,
		valid:      false,
	}
	fmt.Println("PASSENGER NAME: ", ticket1.name)
	fmt.Println("PASSPORTNO: ", ticket1.passportno)

	if ticket1.valid {
		fmt.Println("VALIDATION SUCESSFUL")
	} else {
		fmt.Println("NOT VALID!")
	}

	P := &ticket1.price
	fmt.Println("COST: ", *P)

	*P = *P / 2
	fmt.Println("EMI: ", *P)
	payment := []float64{}
	payment = append(payment, *P)
	payment = append(payment, *P)
	fmt.Println("EMI SCHUDLE")
	totalpaid := 00.00
	for i := 0; i < len(payment); i++ {
		fmt.Println("month", i+1, ";", payment[i])
		totalpaid += payment[i]
	}
	fmt.Println("TOTAL PAID =", totalpaid)
	if totalpaid >= 5645.0 {
		fmt.Println("COMPLETED")
	} else if totalpaid > 0 {
		fmt.Println("1 MONTH BALANCE")
	} else {
		fmt.Println("NOT PAID")
	}
}
