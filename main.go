package main

import (
	"fmt"
	"strings"
)

func main() {
	var ConferenceName = "Go Conference"
    const ConferenceTickets = 50
	var remainingTickets = 50
	var bookings = []string{}

    fmt.Printf("Welcome to %v booking application", ConferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availabe.\n", ConferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
    
	for   {
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
	     
		if userTickets < remainingTickets {
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)
	
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, ConferenceName)
	    
		firstNames := []string{}
		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
  
		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	} else {
		fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		continue
	}
}
}