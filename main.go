package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets type: %T\n, remainingTickets type: %T\n", conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "and these are available:", remainingTickets, "tickets available")
	fmt.Println("Get your tickets here to attend the conference")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	bookings := []string{}

	fmt.Println("Please enter your first name and last name")
	fmt.Scan(&firstName, &lastName)

	fmt.Println("Please enter your email address")
	fmt.Scan(&email)
	//fmt.Println(&remainingTickets)

	fmt.Println("Please enter number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	bookings = append(bookings, firstName+" "+lastName+" "+email)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets remaining for %v", remainingTickets, conferenceName)

	fmt.Printf("These are all the Bookings: %v\n", bookings)
}
