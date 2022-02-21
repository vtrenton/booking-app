package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint8

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter your number of tickets:")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := inputValidation(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %s %s for booking %d Tickets! You will recieve an email at %s.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%d remaining tickets left\n", remainingTickets)
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of our current bookings: %s\n", firstNames)

		} else {
			if !isValidName {
				fmt.Printf("The name %s %s does not contain at least 2 letters in either the first or last name\n", firstName, lastName)
			}
			if !isValidEmail {
				fmt.Printf("The email address %s is malformed\n", email)
			}
			if !isValidTicketNumber {
				fmt.Printf("The ticket value %d is out of range. %d tickets remain\n", userTickets, remainingTickets)
			}
			fmt.Println("invalid input data, please try again")
			//continue
		}
	}
}

func greetUsers(confName string, confTickets uint8, remainTickets uint8) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still availible\n", confTickets, remainTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func inputValidation(firstName string, lastName string, email string, userTickets uint8, remainingTickets uint8) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
