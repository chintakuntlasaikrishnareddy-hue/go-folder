package main

import "fmt"

type Book struct {
	Title     string
	Author    string
	Pages     int
	Available bool
}

// Display book details
func (b Book) ShowDetails() {
	fmt.Println("\n------ Book Details ------")
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Pages:", b.Pages)
}

// Display availability
func (b Book) ShowAvailability() {
	if b.Available {
		fmt.Println("Book is Available")
	} else {
		fmt.Println("Book is Not Available")
	}
}

// Check if it is a big book
func (b Book) IsBigBook() bool {
	return b.Pages > 500
}

func main() {

	var book Book

	// User Input
	fmt.Print("Enter Book Title: ")
	fmt.Scan(&book.Title)

	fmt.Print("Enter Author Name: ")
	fmt.Scan(&book.Author)

	fmt.Print("Enter Number of Pages: ")
	fmt.Scan(&book.Pages)

	fmt.Print("Is the Book Available (true/false): ")
	fmt.Scan(&book.Available)

	// Method Calls
	book.ShowDetails()
	book.ShowAvailability()

	if book.IsBigBook() {
		fmt.Println("Big Book")
	} else {
		fmt.Println("Normal Book")
	}
}
