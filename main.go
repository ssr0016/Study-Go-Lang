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

package main

import "fmt" // fmt stands for Format package

func main() {
	//Variables
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceName, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

}
