package main

import "fmt"

func main() {
	var ConferenceName = "Go Conference"
    const ConferenceTickets = 50
	var remainingTickets = 50

    fmt.Printf("Welcome to %v booking application", ConferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availabe.\n", ConferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
    
	var firstName string
	var lastName string
	var email string
    var userTickets int

	fmt.Println("Enter your first name: ")
    fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)
  
	fmt.Println("Enter number of tickets you want: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

    fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, ConferenceName)
}