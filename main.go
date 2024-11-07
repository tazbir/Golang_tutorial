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

	var userName string
	var ticketQuantity int

	fmt.Scan(&userName)

	ticketQuantity = 2
	fmt.Printf("Hello %v, you have booked %v tickets \n", userName, ticketQuantity)

}
