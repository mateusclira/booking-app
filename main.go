package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets type: %T\n, remainingTickets type: %T\n", conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "and these are available:", remainingTickets, "tickets available")
	fmt.Println("Get your tickets here to attend the conference")
	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Please enter your first name and last name")
		fmt.Scan(&firstName, &lastName)

		fmt.Println("Please enter your email address")
		fmt.Scan(&email)
		//fmt.Println(&remainingTickets)

		fmt.Println("Please enter number of tickets")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName+" "+email)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("We have %v tickets remaining for %v", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all the people who has booked: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Sorry, we have no tickets remaining")
				break
			}
		} else {
			fmt.Printf("Sorry, we have only %v tickets remaining so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}

}
