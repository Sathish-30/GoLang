package helper

import "strings"

// capitalize the function name which will be exported
func CheckValidUser(firstName string, lastName string, emailID string) (bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	isValidEmail := strings.Contains(emailID, "@")

	return isValidName, isValidEmail
}
