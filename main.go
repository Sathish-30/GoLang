package main

// Go progarm are managed by the packages
// A package is a collection of source file
// fmt stands for format package
import (
	"Booking-App/helper"
	"fmt"
	"strconv"
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

	// Slicing of map initialization
	var bookingDetails = make([]map[string]string , 0)

	type UserData struct {
		firstName string
		lastName string
		emailID string
		numberOfTickets uint
	}
	// To append element in the slice we do it by
	//bookingNames = append(bookingNames , "Sathish")

	// To retrieve data from slice we get it by index
	// names[0] ... 
	for {
		if remainingTickets == 0 {
			fmt.Printf("The first Name of the users who booked the tickets : %v\n",printFirstNamesOfPersons(bookingDetails));
			break;
		}

 		firstName , lastName , emailID , userTickets := getUserInput()
		status := false 
		isValidName , isValidEmail := helper.CheckValidUser(firstName , lastName , emailID )
		// isValidName := len(firstName) >= 2 && len(lastName) >= 2

		// isValidEmail := strings.Contains(emailID , "@")
		
		// isValidTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName {
			remainingTickets -= userTickets

			//fmt.Printf("The first name is %#v and the last name is %#v and the email ID is %v\n",firstName,lastName,emailID)
			
			//fmt.Printf("The whole array %v \n",bookingNames)

			//fmt.Printf("The first value %#v\n",bookingNames[0])

			//fmt.Printf("The type of the array %T\n",bookingNames)

			//fmt.Printf("The length of the array %v\n",len(bookingNames))
			if checkTickets(userTickets , remainingTickets){
				fmt.Println("The Tickets limit have been exceeded")
				remainingTickets += userTickets
				fmt.Printf("You can only book seats for limited remaining seats %v\n",remainingTickets)
				status = true;
			}
			if !status{
				var userData = make(map[string]string);
				userData["firstName"] = firstName
				userData["lastName"] = lastName
				userData["emailID"] = emailID
				// The below line of code will convert the uint to a string
				// strconv is a package which help in conversion of variables to strings
				userData["noOfTickets"] = strconv.FormatUint(uint64(userTickets),10)
				bookingDetails = append(bookingDetails , userData)
				fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n",firstName,lastName,userTickets,emailID)
			}

			
			
			// noTicketRemaining := remainingTickets == 0
			// The above comment will assign whether a true or false

		
		}else{
			
			if !(isValidEmail && isValidName){
				fmt.Println("There is a error in the given person detail ")
			}
		}
	}
}  

func printFirstNamesOfPersons(bookingDetails []map[string]string) []string {
	firstNamesOfPersons := []string{}
	for _ , booking :=   range bookingDetails{
		firstNamesOfPersons = append(firstNamesOfPersons, booking["firstName"])
	}	
	return firstNamesOfPersons;
}



func checkTickets(userTickets uint , remainingTickets uint) bool{
	isValidTickets := remainingTickets > 0 && userTickets <= remainingTickets
	return isValidTickets
}

func getUserInput() (string , string , string , uint){
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
		return firstName , lastName , emailID , userTickets;
}