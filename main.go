package main

import (
	helper "booking-app/utils"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50

//var bookings = make([]map[string]string, 0)

var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1) //amount of go routines
		go sendTicket(firstName, lastName, userTickets, email)

		firstNames := getFirstNames()
		fmt.Printf("These are all the people who has booked: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("Sorry, we have no tickets remaining")
			//break
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "and these are available:", remainingTickets, "tickets available")
	fmt.Println("Get your tickets here to attend the conference")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//creating it with struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// How to do it with mapping

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["tickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets remaining for %v", remainingTickets, conferenceName)
}

func sendTicket(firstName string, lastName string, userTickets uint, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#############################################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("#############################################")
	wg.Done()
}
