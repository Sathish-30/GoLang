package main

// Go progarm are managed by the packages
// A package is a collection of source file
// fmt stands for format package
import (
	"fmt"
	"strings"
)

// It requires a entry point , mean a main function
func main(){
	// If we declare a varibale but not used the variable it will throw a error

	// In const keyword we can't change the value 
	
	// Print format function is used to print using format specifier

	/*

	Switch Case

	var cityName string
	fmt.Print("Enter the city name : ");
	fmt.Scan(&cityName);

	switch cityName{
		case "New York":
		// execute code for booking new york conference tickets
		case "Singapore":
		// execute code for booking new york conference tickets
		case "London":
		// execute code for booking new york conference tickets
		case "Hong Kong":
		// execute code for booking new york conference tickets
		default:
			fmt.Println("No valid city Selected");
		// execute code when no given cites takes place
	}

	*/

	var remainingTickets uint
	const totalTickets uint = 10
	remainingTickets = totalTickets

	// It is for Initializing a static array element
	
	//var bookings = [10]string{"Sathish","sandi","karthi","chittappa","ratish"}

	// Array Initialization
	// Arrays are index based 

	// Slicing initialization
	var bookingNames []string

	// To append element in the slice we do it by
	//bookingNames = append(bookingNames , "Sathish")

	// To retrieve data from slice we get it by index
	// names[0] ... 
	for {
		var firstName string
		var lastName string
		var emailID string
		var userTickets uint

		fmt.Print("Enter the first name : ")
		fmt.Scan(&firstName)

		fmt.Print("Enter the last name : ")
		fmt.Scan(&lastName)

		fmt.Print("Enter the mail Id : ");
		fmt.Scan(&emailID);

		fmt.Print("Enter the user tickets : ");
		fmt.Scan(&userTickets);

		// isValidName := len(firstName) >= 2 && len(lastName) >= 2

		// isValidEmail := strings.Contains(emailID , "@")
		
		// isValidTickets := userTickets > 0 && userTickets <= remainingTickets

		if checkValidUser(firstName , lastName , emailID , userTickets , remainingTickets){
			bookingNames = append( bookingNames, firstName+" "+lastName)

			remainingTickets -= userTickets

			//fmt.Printf("The first name is %#v and the last name is %#v and the email ID is %v\n",firstName,lastName,emailID)
			
			fmt.Printf("The whole array %v \n",bookingNames)

			fmt.Printf("The first value %#v\n",bookingNames[0])

			fmt.Printf("The type of the array %T\n",bookingNames)

			fmt.Printf("The length of the array %v\n",len(bookingNames))

			fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n",firstName,lastName,userTickets,emailID)
			firstNamesOfPersons := []string{}

			for _ , name := range bookingNames{
				var names = strings.Fields(name)
				firstNamesOfPersons = append(firstNamesOfPersons, names[0])
			}
			
			// noTicketRemaining := remainingTickets == 0
			// The above comment will assign whether a true or false

			if remainingTickets == 0{
				fmt.Printf("The first Name of the users who booked the tickets : %v\n",firstNamesOfPersons);
				fmt.Print("The Conference has been ended");
				break;
			}
		}else{
			if !checkValidUser(firstName , lastName , emailID , userTickets , remainingTickets) {
				fmt.Println("First name and the last name you entered is too short")
			}
			fmt.Printf("We only have %v tickets , so you can't book %v tickets",remainingTickets,userTickets);
		}
	}
}

func checkValidUser (firstName string, lastName string, emailID string , userTickets uint , remainingTickets uint){
		isValidName := len(firstName) >= 2 && len(lastName) >= 2

		isValidEmail := strings.Contains(emailID , "@")
		
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets

		return isValidTickets && isValidName && isValidEmail
}