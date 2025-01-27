package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var availableTickets uint = 50
	var bookings []string

	greetUsers(conferenceName, conferenceTickets, availableTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketAmount := validateUserInput(firstName, lastName, email, userTickets, availableTickets)

		if isValidName && isValidEmail && isValidTicketAmount{
			availableTickets = availableTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Slice type: %T\n", bookings[0])
			fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", availableTickets, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if availableTickets == 0{
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		}else{
			//fmt.Printf("We only have %v tickets remaining. You can't book %v tickets\n", availableTickets, userTickets)
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketAmount{
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers(confName string, confTickets int, availableTickets uint){
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets with %v still available.\n", confTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string{
	firstNames := []string{}
		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, availableTickets uint) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketAmount := userTickets > 0 && userTickets < availableTickets

	return isValidName, isValidEmail, isValidTicketAmount
}