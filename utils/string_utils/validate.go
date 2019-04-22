package string_utils

import "regexp"

func IsValidEmail(potentialEmail string) bool {
	emailPattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailPattern.MatchString(potentialEmail)
}

func IsValidPassword(potentialPassword string) bool {
	minPasswordLength := 8
	return len(potentialPassword) >= minPasswordLength
}
