package main

import (
	"fmt"
	"strings"
)

func main() {
	confName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T and confName is %T\n", conferenceTickets, remainingTickets, confName)

	// fmt.Printf("Welcome to the %v book application\n", confName)
	// fmt.Printf("We have a total %v tickets and %v tickets remaining\n", conferenceTickets, remainingTickets)
	// fmt.Println("Get your tickets here")

	greetUser(confName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want?")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNum := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNum {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The user %v %v bought %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, confName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our event is booked out. Come back next year :)")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Your email doesn't contain @")
			}
			if !isValidTicketNum {
				fmt.Println("The ticket quantity is invalid")
			}
			fmt.Println("Please, try again.")
		}
	}
}

func greetUser(confName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to the %v booking application\n", confName)
	fmt.Printf("We have a total %v tickets and %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")
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
	isValidTicketNum := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNum
}
