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
	"fmt"
	"strings"
)

func main() {
	//Variables
	// var conferenceName string = "Go conference"
	// const conferenceTickets int = 50
	// var remainingTickets int = 50

	//There is alternative systax for assigning variables
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	//Arrays & slices
	//Size & types
	// var bookings []string
	//alternative types for slice
	bookings := []string{}

	//Note: %v is a placeholder
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//loops
	for { //infinite loops
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

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			//if this is true we need to end the program
			// it will break the loop
			// break
			//continue statement causes to skip remainder of its body
			continue // it means continue the next iteration in the loop

			//Book ticket logic
			remainingTickets = remainingTickets - userTickets
			//slices
			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("The whole slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)        //knowing the type
			// fmt.Printf("Slice length: %v\n", len(bookings)) // knowing the length

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			//_ is a blan identifier
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				//slice
				firstNames = append(firstNames, names[0])
			}

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

		}

	}
}
