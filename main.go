package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings []string

func main() {
	// fmt.Printf("ConferenceTickets is of type %T and remainingTickets is of type %T \n", conferenceTickets, remainingTickets)

	greetUser()

	for {
		firstName, lastName, email, ticketQuantity := getUserInput()

		isValidName, isValidEmail, isValidTicketQuantity := validateUserInput(firstName, lastName, email, ticketQuantity, remainingTickets)

		if isValidName && isValidEmail && isValidTicketQuantity {

			bookTicket(remainingTickets, ticketQuantity, firstName, lastName, bookings, conferenceName, email)

			firstNames := getFirstName(bookings)
			fmt.Println("All bookings so far: ", firstNames)

			var noTicketsRemaining bool = remainingTickets == 0

			if noTicketsRemaining {
				fmt.Println("All tickets have been booked. No more tickets available.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name. Name should be at least 2 characters long.")
			}
			if !isValidEmail {
				fmt.Println("Invalid email. Email should contain @.")
			}
			if !isValidTicketQuantity {
				fmt.Println("Invalid ticket quantity. Please enter a valid ticket quantity.")
			}
			fmt.Println("Invalid input. Please try again.")
		}

	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have %v tickets remaining and %v tickets in total \n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend the conference")

}

func getFirstName(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, ticketQuantity uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketQuantity := ticketQuantity > 0 && ticketQuantity <= remainingTickets

	return isValidName, isValidEmail, isValidTicketQuantity
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var ticketQuantity uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to book:")
	fmt.Scan(&ticketQuantity)

	return firstName, lastName, email, ticketQuantity
}

func bookTicket(remainingTickets uint, ticketQuantity uint, firstName string, lastName string, conferenceName string, email string) {
	remainingTickets = remainingTickets - ticketQuantity

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Congratulations %v %v, you have successfully booked %v tickets for the %v conference. \n", firstName, lastName, ticketQuantity, conferenceName)
	fmt.Printf("An email has been sent to %v with the details of your booking. \n", email)
	fmt.Printf("There are %v tickets remaining for the %v conference. \n", remainingTickets, conferenceName)

}
