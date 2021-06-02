package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	d := Developer{
		Individual: Employee{
			Id:        1,
			FirstName: "Duc",
			LastName:  "Do",
		},
		HourlyRate: 10,
	}

	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	fmt.Println("Hours worked on Monday: ", d.WorkWeek[Monday])
	fmt.Println("Hours worked on Tuesday: ", d.WorkWeek[Tuesday])
	fmt.Println("Hours worked on this week: ", d.HoursWorked())
}

func (d *Developer) LogHours(day Weekday, hour int) {
	if day < 0 || day > 6 {
		fmt.Println("Invaid day")
		return
	}
	d.WorkWeek[day] = hour
}

func (d *Developer) HoursWorked() int {
	sum := 0
	for _, v := range d.WorkWeek {
		sum += v
	}
	return sum
}
