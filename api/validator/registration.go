package validator

import (
	"net/mail"
	"regexp"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidPhoneNumber(phoneNumber string) bool {
	pattern := `(^\+?[0-9]{9,15})$`

	match, err := regexp.MatchString(pattern, phoneNumber)
	if err != nil {
		return false
	}

	return match
}
