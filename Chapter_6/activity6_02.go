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
	pay, err := PayWeeks(80, 50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Pay for this week: ", pay)
	}

	pay, err = PayWeeks(30, 120)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Pay for this week: ", pay)
	}
}

func PayWeeks(workHours, HourlyRate int) (int, error) {
	if workHours <= 0 || workHours > 80 {
		return 0, ErrHourWorked
	}

	if HourlyRate < 10 || HourlyRate > 75 {
		return 0, ErrHourlyRate
	}

	if workHours <= 40 {
		return workHours * HourlyRate, nil
	} else {
		return 40*HourlyRate + (workHours-40)*2*HourlyRate, nil
	}
}
