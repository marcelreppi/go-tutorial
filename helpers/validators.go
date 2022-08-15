package helpers

import "strings"

func IsValidName(name string) bool {
	return len(name) > 2
}

func IsValidEmail(email string) bool {
	return strings.Contains(email, "@")
}

func IsValidTicketNumber(numberOfTickets uint, remainingTickets uint) bool {
	return numberOfTickets > 0 && numberOfTickets <= remainingTickets
}
