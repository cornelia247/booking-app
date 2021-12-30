package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Super Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	// unlimited loop

	firstName, lastName, email, userTickets := getUserInput()
	isValidEmail, isValidName, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
	if isValidEmail && isValidName && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		//for each loop
		firstNames := getFirstNames()
		fmt.Printf("These are the First names that booked for %v: %v\n", conferenceName, firstNames)
		// if else loop
		if remainingTickets <= 0 {
			// end Program
			fmt.Printf("%v is fully booked!  \n", conferenceName)
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered is wrong")
		}
		if !isValidTicketNumber {
			fmt.Printf("The number of ticket you entered is invalid there are %v tickects remaining\n", remainingTickets)
		}
	}

	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking Application\n", conferenceName)
	fmt.Printf("We have Total of %v Tickets and %v Tickets are still Availablle\n", conferenceTickets, remainingTickets)
	fmt.Println("Book Your Attendance ticket here!")
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter your  Number of Tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//map for a user
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a Confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining for %v \n", remainingTickets, conferenceName)
}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done()
}
