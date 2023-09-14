package main

import "fmt"

func main() {
	confName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T and confName is %T\n", conferenceTickets, remainingTickets, confName)

	fmt.Printf("Welcome to the %v book application\n", confName)
	fmt.Printf("We have a total %v tickets and %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

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

	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The user %v %v bought %v tickets, you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, confName)

	fmt.Printf("These are all our bookings: %v", bookings)
}
