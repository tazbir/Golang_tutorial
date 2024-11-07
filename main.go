package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	// fmt.Printf("ConferenceTickets is of type %T and remainingTickets is of type %T \n", conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have %v tickets remaining and %v tickets in total \n", remainingTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend the conference")

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

	remainingTickets = conferenceTickets - ticketQuantity

	fmt.Printf("Congratulations %v %v, you have successfully booked %v tickets for the %v conference. \n", firstName, lastName, ticketQuantity, conferenceName)
	fmt.Printf("An email has been sent to %v with the details of your booking. \n", email)
	fmt.Printf("There are %v tickets remaining for the %v conference. \n", remainingTickets, conferenceName)
}
