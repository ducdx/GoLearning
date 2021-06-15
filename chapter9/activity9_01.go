package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrInvalidSSNLength   = errors.New("invalid SSN length")
	ErrInvalidSSNNumbers  = errors.New("invalid SSN numbers")
	ErrInvalidSSNPrefix   = errors.New("invalid SSN prefix")
	ErrInvalidDigitPlaces = errors.New("invalid digits place")
)

func validateSSNlength(ssn string) (string, error) {
	if len(ssn) != 9 {
		return ssn, ErrInvalidSSNLength
	}
	return ssn, nil
}

func main() {
	str := "012-8-678"
	_, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
}
