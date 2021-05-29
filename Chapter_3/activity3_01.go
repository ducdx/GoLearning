package main

import (
	"fmt"
	"unicode"
)

func main() {
	pwd := "ThisisAPassowrd"
	validatePassword(pwd)

	pwd = "Aa@123456"
	validatePassword(pwd)

	pwd = "aaaaaaaaaaaaaaaaaaaaaaaa"
	validatePassword(pwd)
}

func validatePassword(pwd string) {
	valid := checkPassword(pwd)
	if valid {
		fmt.Println("This is a good password")
	} else {
		fmt.Println("This is wrong format password")
	}
}

func checkPassword(pwd string) bool {
	pwR := []rune(pwd)
	if len(pwR) < 8 || len(pwR) > 15 {
		return false
	}
	hasUpper, hasLower, hasNumber, hasSymbol := false, false, false, false
	for _, val := range pwR {
		if unicode.IsUpper(val) {
			hasUpper = true
		}

		if unicode.IsLower(val) {
			hasLower = true
		}

		if unicode.IsDigit(val) {
			hasNumber = true
		}

		if unicode.IsPunct(val) || unicode.IsSymbol(val) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}
