package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate = errors.New("Invalid Hourly Rate")
	ErrHourWorked = errors.New("Invalid Hour Worked")
)

func main() {
	pay := PayWeeks(80, 50)
	fmt.Println("Pay for this week 1 - : ", pay)

	pay = PayWeeks(30, 120)
	fmt.Println("Pay for this week 2 - : ", pay)

	pay = PayWeeks(90, 50)
	fmt.Println("Pay for this week 3 - : ", pay)

	pay = PayWeeks(90, 120)
	fmt.Println("Pay for this week 4 - : ", pay)
}

func PayWeeks(workHours, HourlyRate int) int {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrHourWorked {
				fmt.Printf("\n\nHour Worked: %d\nError: %v\n", workHours, r)
			} else if r == ErrHourlyRate {
				fmt.Printf("\n\nHourly Rate: %d\nError: %v\n", HourlyRate, r)
			}
		}
	}()

	if workHours <= 0 || workHours > 80 {
		panic(ErrHourWorked)
	}

	if HourlyRate < 10 || HourlyRate > 75 {
		panic(ErrHourlyRate)
	}

	if workHours <= 40 {
		return workHours * HourlyRate
	} else {
		return 40*HourlyRate + (workHours-40)*2*HourlyRate
	}
}
