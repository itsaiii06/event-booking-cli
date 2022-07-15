package common

import (
	"strings"
)

func ValidateInput(userName string, userEmail string, userTickets uint, MAX_TICKETS uint) bool {
	var isValidName bool = len(userName) >= 2
	var isValidEmailName bool = len(userName) >= 2 && strings.Contains(userEmail, "@")
	var isValidTickets bool = userTickets < MAX_TICKETS && userTickets > 0

	return isValidName && isValidEmailName && isValidTickets

}
