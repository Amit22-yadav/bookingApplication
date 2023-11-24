package main

import (
	"fmt"
	"study-go/helper"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

// creating structs in go lang

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	// calling the functions to greet the users.
	greetUsers()
	// fmt.Println("Welcome to" ,conferenceName ,"booking Application")
	for {
		// Geeting the input from users
		firstName, lastName, email, userTickets := getUserInput()

		// validating the user input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			// calling the functions to print the firstNames.
			firstNames := printFristNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

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

func greetUsers() {
	fmt.Printf("Welcome to %v Booking application!!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!!!\n ", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!!!")

}
func printFristNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)

		// firstNames = append(firstNames, names[0])
		firstNames = append(firstNames, booking.firstName)
	}
	// fmt.Printf("These are all our bookings: %v \n", bookings)
	return firstNames

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	// fmt.Printf("The whole Slice: %v\n", bookings)
	// fmt.Printf("The first value is : %v\n", bookings[0])

	fmt.Printf("Thank you  %v  %v for booking %v tickets.You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v \n", remainingTickets, conferenceName)
}
