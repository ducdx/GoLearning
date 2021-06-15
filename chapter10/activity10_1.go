package main

import (
	"fmt"
	"time"
)

func main() {
	Day := time.Now().Weekday()
	Hour := time.Now().Hour()
	fmt.Println("Day: ", Day, "--- Hour: ", Hour)
	if Day.String() == "Monday" {
		if Hour >= 1 {
			fmt.Println("Performing full blown test!")
		} else {
			fmt.Println("Performing hit-n-run test!")
		}
	} else {
		fmt.Println("Performing hit-n-run test!")
	}
}
