package utils

import (
	"net/mail"
	"regexp"
)

func ValidateString(value string, minLength int, maxLength int) bool {
	n := len(value)
	if n < minLength || n > maxLength {
		return false
	}
	return true
}

func IsValidURL(input string) bool {
	regexPattern := `^(http|https):\/\/[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)+([\/?].*)?$`

	re := regexp.MustCompile(regexPattern)
	return re.MatchString(input)
}

func IsValidEmail(value string) bool {
	if _, err := mail.ParseAddress(value); err != nil {
		return false
	}
	return true
}

func IsValidName(value string) bool {
	if !ValidateString(value, 5, 100) {
		return false
	}
	if !regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(value) {
		return false
	}
	return true
}

func IsValidPassword(value string) bool {
	return ValidateString(value, 6, 100)
}
