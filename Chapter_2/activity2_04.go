package main

import (
	"fmt"
	"errors"
)

func main() {
	x := 21
	if err := validateNumber(&x); err != nil {
		fmt.Println(err)
	}
}

func validateNumber(a *int) error {
	if *a < 0 {
		return errors.New("Input can't be negative")
	} else if *a > 100 {
		return errors.New("Input can't be greater than 100")
	} else if *a%7 == 0 {
		return errors.New("Input can't divisible by 7")
	} else {
		return nil
	}
}
