package main

import (
	"fmt"
	"time"
)

func checkstatus(serviceName string) {
	fmt.Printf("starting the monitering process for%s.....\n", serviceName)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("service %sis Healthy\n", serviceName)
}

func main() {
	go checkstatus("AuthAPI")
	go checkstatus("payment Gateway")
	go checkstatus("DataBase Cluster")
	time.Sleep(1 * time.Second)
	fmt.Println("All initial checks executes")
}
