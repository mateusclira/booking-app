package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, firstName, lastName, email, conferenceName, bookings)

			firstNames := getFirstNames(bookings)
			fmt.Printf("These are all the people who has booked: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Sorry, we have no tickets remaining")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name is invalid, try again")
			}
			if !isValidEmail {
				fmt.Println("Your email is invalid, try again")
			}
			if !isValidTicketNumber {
				fmt.Println("The amount you are trying to buy is invalid, try again")
			}
		}
	}
}

func greetUsers(confName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Println("We have a total of", conferenceTickets, "and these are available:", remainingTickets, "tickets available")
	fmt.Println("Get your tickets here to attend the conference")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, email string, conferenceName string, bookings []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName+" "+email)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets remaining for %v", remainingTickets, conferenceName)
}
