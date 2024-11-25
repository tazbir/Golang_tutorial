package main

import "strings"


func validateUserInput(firstName string, lastName string, email string, ticketQuantity uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketQuantity := ticketQuantity > 0 && ticketQuantity <= remainingTickets

	return isValidName, isValidEmail, isValidTicketQuantity
}