package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

	var bookings [50]string
	bookings[0] = "Antony"

	var userName string
	var userTickets int
	// ask for user name
	fmt.Print("Enter your first name:")
	fmt.Scan(&userName)
	// Pointers

	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
