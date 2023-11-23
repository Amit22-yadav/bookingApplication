package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	var bookings []string

	// fmt.Println("Welcome to" ,conferenceName ,"booking Application")
	fmt.Printf("Welcome to %v booking Application\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n ", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets int
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your Email: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		// userName = "Tom"
		// userTickets = 2
		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The whole Slice: %v\n", bookings)
			// fmt.Printf("The first value is : %v\n", bookings[0])

			fmt.Printf("Thank you  %v  %v for booking %v tickets.You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are remaining for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			fmt.Printf("These are all our bookings: %v \n", bookings)

			if remainingTickets == 0 {
				//end program
				fmt.Println("All tickets are booked")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("First name or last name you entred is too short !!")
			}
			if !isValidEmail {
				fmt.Println("Email address you entred dosen't contains @ sign!!!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entred is incorrect!!")
			}

			// fmt.Printf("We only have %v tickets remaining,so you can't book %v tickets\n", remainingTickets, userTickets)
		}

	}

}

// User input validation
