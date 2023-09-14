package main

import "fmt"

func main() {
	var confName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T and confName is %T\n", conferenceTickets, remainingTickets, confName)

	fmt.Printf("Welcome to the %v book application\n", confName)
	fmt.Printf("We have a total %v tickets and %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

	var userName string
	var userTickets int

	userName = "Teste"
	userTickets = 1

	fmt.Printf("The user %v bought %v tickets\n", userName, userTickets)
}
