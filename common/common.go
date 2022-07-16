package common

import (
	"fmt"
	"strings"
)

func ValidateInput(userName string, userEmail string, userTickets uint, MAX_TICKETS uint) bool {
	var isValidName bool = len(userName) >= 2
	var isValidEmailName bool = len(userName) >= 2 && strings.Contains(userEmail, "@")
	var isValidTickets bool = userTickets <= MAX_TICKETS && userTickets > 0

	// error handling
	if !isValidName {
		fmt.Println("Check your name and try again ")
	}
	if !isValidEmailName {
		fmt.Println("Check your email address and try again ")
	}
	if !isValidTickets {
		fmt.Printf("Sorry we don't have that many tickets available right now \n")
	}

	return isValidName && isValidEmailName && isValidTickets

}

func GetFirstName(fullName string) string {
	var firstName = strings.Fields(fullName)

	return firstName[0]

}
