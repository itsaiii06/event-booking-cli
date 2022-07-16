package main

import (
	"event-booking-cli/common"
	"fmt"
	"sync"
	"time"
)

const TOTAL_TICKETS uint = 100

var remainingTickets = TOTAL_TICKETS
var attendees = make([]UserData, 0) // zero wont matter becuase it will increase anyways
var firstNames = []string{}

type UserData struct {
	userName    string
	userEmail   string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	getGreeting()

	// for remainingTickets > 0 {
	userName, userEmail, userTickets := getInput() //gets the input

	if !common.ValidateInput(userName, userEmail, userTickets, remainingTickets) {
		return
		// continue // not true so skip the portion below
	}

	if remainingTickets == 0 {
		fmt.Println("")
		fmt.Println("we're all out !")
		fmt.Println("*********** TRY NEXT YEAR ***********")

	}

	var userData UserData = getTicket(userName, userEmail, userTickets)

	wg.Add(1)
	go sendTicket(userData)

	var firstName string = common.GetFirstName(userName)
	firstNames = append(firstNames, firstName)

	// }

	wg.Wait()

}

func getGreeting() {
	fmt.Println("******** GOPHER CON TICKET BOOKING CLI ********")
	fmt.Println("")

	fmt.Println("Kindly enter the following details to begin the registration process : ")
}

func getInput() (string, string, uint) {
	var userName string
	var userName2 string
	var userEmail string
	var userTickets uint

	fmt.Println("Enter your name : ")
	fmt.Scan(&userName)
	fmt.Scan(&userName2)

	fmt.Println("Enter your email address : ")
	fmt.Scan(&userEmail)

	fmt.Println("How many tickets do you want to book? ")
	fmt.Scan(&userTickets)

	fmt.Println("")

	return userName + " " + userName2, userEmail, userTickets

}

func getTicket(userName string, userEmail string, userTickets uint) UserData {
	remainingTickets -= userTickets

	var userData = UserData{
		userName:    userName,
		userEmail:   userEmail,
		userTickets: userTickets,
	}

	attendees = append(attendees, userData)

	// fmt.Printf("%v tickets registered for %v \n", userTickets, userName)
	// fmt.Printf("attendees = %v \n\n", attendees)

	return userData

}

func sendTicket(data UserData) {
	fmt.Println("GENERATING TICKET .....")
	time.Sleep(10 * time.Second) // sleep stops the execution for the current thread
	fmt.Println("################################################")
	fmt.Printf("%v tickets have been booked for %v \n", data.userTickets, data.userName)
	fmt.Println("Check your email : ", data.userEmail)
	fmt.Println("")
	fmt.Println("*********** SEE YOU AT THE GOPHER CON ***********")
	fmt.Println("################################################")
	fmt.Println("")

	wg.Done()

}
