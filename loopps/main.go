package main

import "fmt"

func main() {
	/*	// standard counter based loop

			//FOR loop

			for s := 1; s <= 7; s++ {
				fmt.Println("itration: ", s)
			}
		}

		// WHILE LOOP
		/* count := 1
		for count <= 100 {
			fmt.Println("count is:", count)

			count++
		}*/

	// infinite loop

	iteration := 0
	for {
		iteration++
		if iteration > 3 {
			fmt.Println("Breaking out of infinite loop")
			break
		}
		fmt.Println("Running contineously")
	}

	service := []string{"Auth API", "Payment Gateway", "DataBase Watcher"}

	for index, name := range service {
		fmt.Printf("Index:%d|serviceName:%s\n", index, name)
	}
}
