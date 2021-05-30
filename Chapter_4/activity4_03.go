package main

import "fmt"

func main() {
	weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	weekdays = append(weekdays[6:], weekdays[:6]...)
	fmt.Println("Weekday: ", weekdays)
}
