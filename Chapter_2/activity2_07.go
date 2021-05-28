package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Monday

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born day is a week day")
	case time.Saturday, time.Sunday:
		fmt.Println("Born day is a weekend day")
	default:
		fmt.Println("Error, day born not valid")

	}
}
