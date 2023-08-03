package main

// Go progarm are managed by the packages
// A package is a collection of source file
import "fmt"

// It requires a entry point , mean a main function
func main(){
	// If we declare a varibale but not used the variable it will throw a error
	var conferenceName = "Go Conference"
	// In const keyword we can't change the value 
	const conferenceTickets = 50
	var remainingTickets = 30

	// Print format function is used to print using format specifier
	fmt.Printf("Welcome to %v booking application\n",conferenceName)

	fmt.Printf("The total conference Tickets : %v\n",conferenceTickets)

	fmt.Printf("Remaining Tickets : %v", remainingTickets )
}

