//go mod init <module path>
//ex.go mod init booking-app
//Note:
//>Creates a new module
// >Module path can correspond to a repository you plan to
// publish your module to (e.g. github/nana/booking-app)

// Note:
//All our code must belong to a package
//The first statement in Go file must be "package..."

//To run the application
//go run main.go or the name of an application
//Print statement > fmt.Println("")

//%T is for checking which types of variables
// fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

package main

// fmt stands for Format package
import (
	"booking-app/helper"
	"fmt"
	// "strconv"
)

//Variables
// var conferenceName string = "Go conference"
// const conferenceTickets int = 50
// var remainingTickets int = 50

// There is alternative systax for assigning variables
var conferenceName = "Go conference"

const conferenceTickets int = 50

//export variable
var remainingTickets uint = 50

// Arrays & slices
// Size & types
// var bookings []string
// alternative types for slice

//Map
// var bookings = make([]map[string]string, 0)

// Struct
var bookings = make([]UserData, 0)

// Struct - it can hold mixed data types
type UserData struct {
	//block
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	//Greetings
	//Paramenters > conferenceName, conferenceTickets, remainingTickets
	greetUsers()

	//loops
	for { //infinite loops

		//Get user input... function
		firstName, lastName, email, userTickets := getUserInput()

		//Validate the user input using a function
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			//function for Booking
			bookTicket(userTickets, firstName, lastName, email)
			sendTicket(userTickets, firstName, lastName, email)

			//call function get & print first names
			firstNames := getFirstName()
			fmt.Printf("The first names of bookings are:  %v\n", firstNames)

			//check if there is an available tickets... if else statement
			// var noTicketsRemaining bool = remainingTickets == 0
			//alternative systax
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end program logic
				fmt.Println("Our conference is booked out. Come back next year.")
				//break is breaking the loop
				break

			}

		} else {
			if !isValidName {
				fmt.Println("firt name or last you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			//if this is true we need to end the program
			// it will break the loop
			// break
			//continue statement causes to skip remainder of its body
			// continue // it means continue the next iteration in the loop

		}
	}

}

// the function must be called insede the main function because the main function is the entry point
func greetUsers() {
	//Note: %v is a placeholder
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

// fucntion for printing the first name
// booking parameter is a slice of strings
func getFirstName() []string {
	firstNames := []string{}
	//_ is a blank identifier
	for _, booking := range bookings {
		//slice
		firstNames = append(firstNames, booking.firstName)
	}
	//returning values from a function
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	//Data types... string, integer, boolean, maps, arrays ...
	//Defining the variable types
	//Ask user for their name
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//Ask the user
	//& is a pointer... memory address
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	//note you should to return all the initial values or input values or  data back to the main

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	//Book ticket logic
	remainingTickets = remainingTickets - userTickets

	//create a map for each or a user
	// types of map[key]values... key and value
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// //srtconv, string convertion
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//slices
	bookings = append(bookings, userData) // we have all information of user that has key and values pair
	fmt.Printf("List of bookings is %v\n", bookings)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)        //knowing the type
	// fmt.Printf("Slice length: %v\n", len(bookings)) // knowing the length

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// simulate sending ticket on email address
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################################")
}

// //Switch Statement
// city := "London"

// switch city {
// case "New York":
// 	//execute code for booking New York conference tickets
// case "Singapore", "Hong Kong":
// 	//execute code for booking Singapore & Hong Kong conference tickets
// case "London", "Berlin":
// 	//some code here for london and berlin
// case "Mexico City":
// 	//some code here

// default:
// 	fmt.Printf("No valid city selected")
// }
