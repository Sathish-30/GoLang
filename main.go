package main

// Go progarm are managed by the packages
// A package is a collection of source file
// fmt stands for format package
import "fmt"

// It requires a entry point , mean a main function
func main(){
	// If we declare a varibale but not used the variable it will throw a error

	// In const keyword we can't change the value 
	
	// Print format function is used to print using format specifier

	var firstName string
	var lastName string
	var emailID string
	var userTickets uint
	var remainingTickets uint

	const totalTickets uint = 50
	remainingTickets = totalTickets

	// It is for Initializing a static array element
	
	//var bookings = [10]string{"Sathish","sandi","karthi","chittappa","ratish"}

	// Array Initialization
	// Arrays are index based 
	var bookingNames [10]string

	fmt.Print("Enter the first name : ")
	fmt.Scan(&firstName)

	fmt.Print("Enter the last name : ")
	fmt.Scan(&lastName)

	fmt.Print("Enter the mail Id : ")
	fmt.Scan(&emailID)

	fmt.Print("Enter the user tickets : ")
	fmt.Scan(&userTickets)

	bookingNames[0] = firstName+" "+lastName

	remainingTickets -= userTickets

	//fmt.Printf("The first name is %#v and the last name is %#v and the email ID is %v\n",firstName,lastName,emailID)
	
	fmt.Printf("The whole array %v \n",bookingNames)

	fmt.Printf("The first value %#v\n",bookingNames[0])

	fmt.Printf("The type of the array %T\n",bookingNames)

	fmt.Printf("The length of the array %v\n",len(bookingNames))

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n",firstName,lastName,userTickets,emailID)

	//fmt.Printf("%v Tickets are remaining for the conference" , remainingTickets)
	//fmt.Print("The Address of the mail Id : " , &emailID)


}

