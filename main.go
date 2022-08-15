package main

import (
	"booking-app/helpers"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50
const conferenceName = "Go Conference"

var remainingTickets uint = conferenceTickets
var bookings = []Booking{}

type Booking struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for remainingTickets > 0 {

		firstName, lastName, email, numberOfTickets := getUserInput()

		isValidUserInput := validateUserInput(firstName, lastName, email, numberOfTickets)
		if !isValidUserInput {
			continue
		}

		bookTicket(numberOfTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(numberOfTickets, firstName, lastName, email)

		fmt.Printf("These are all our bookings: %v\n\n", getFirstNames())
	}

	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
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
	var numberOfTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&numberOfTickets)

	return firstName, lastName, email, numberOfTickets
}

func validateUserInput(firstName string, lastName string, email string, numberOfTickets uint) bool {
	isValidName := helpers.IsValidName(firstName) && helpers.IsValidName(lastName)
	isValidEmail := helpers.IsValidEmail(email)
	isValidTicketNumber := helpers.IsValidEmail(email)

	isValidUserInput := isValidName && isValidEmail && isValidTicketNumber

	if !isValidUserInput {
		if !isValidName {
			fmt.Println("First or last name is too short")
		}

		if !isValidEmail {
			fmt.Println("Email address does not contain the @ sign.")
		}

		if !isValidTicketNumber {
			fmt.Println("Number of tickets is invalid.")
		}

		fmt.Printf("Your input data is invalid! Try again.\n")
	}

	return isValidUserInput
}

func bookTicket(numberOfTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= numberOfTickets

	booking := Booking{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: numberOfTickets,
	}

	bookings = append(bookings, booking)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, numberOfTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(numberOfTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", numberOfTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done()
}
