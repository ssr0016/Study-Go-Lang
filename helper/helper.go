package helper

import "strings"

// input values in the first parenthesis, input paramenters... the second one a list of output parameter types
// In order to export, Capitalize the first letter... so that it make it available to other packages
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	//user input validation
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	//returning multiole values, so that it will be available in the main fuction
	return isValidName, isValidEmail, isValidTicketNumber // note after returning the values you need to define the types outside the function parenthesis...

}
