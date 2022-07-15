package main

import (
	"event-booking-cli/common"
	"fmt"
)

const TOTAL_TICKETS uint = 100

var remainingTickets = TOTAL_TICKETS
var attendees = []string{}

func main() {
	getGreeting()

	for remainingTickets > 0 {
		userName, userEmail, userTickets := getInput() //gets the input

		if !common.ValidateInput(userName, userEmail, userTickets, TOTAL_TICKETS) {
			fmt.Println("INVALID INPUT : Please try again!")
			fmt.Println("------------------------------------------")
			continue // not true so skip the portion below
		}

		fmt.Println("Correct Input !")
	}

}

func getGreeting() {
	fmt.Println("******** GOPHER CON TICKET BOOKING CLI ********")
	fmt.Println("")

	fmt.Println("Kindly enter the following details to begin the registration process : ")
}

func getInput() (string, string, uint) {
	var userName string
	var userEmail string
	var userTickets uint

	fmt.Println("Enter your name : ")
	fmt.Scan(&userName)

	fmt.Println("Enter your email address : ")
	fmt.Scan(&userEmail)

	fmt.Println("How many tickets do you want to register?: ")
	fmt.Scan(&userTickets)

	fmt.Println("")

	return userName, userEmail, userTickets

}
