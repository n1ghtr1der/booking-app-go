package main

import "fmt"

func main() {
	var confName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Println("Welcome to the", confName, "book application.")
	fmt.Println("We have a total", conferenceTickets, "tickets and", remainingTickets, "tickets remaining")
	fmt.Println("Get your tickets here")
}
