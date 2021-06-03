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

var weekday = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func main() {

	d := Developer{
		Individual: Employee{
			Id:        1,
			FirstName: "Duc",
			LastName:  "Do",
		},
		HourlyRate: 10,
	}

	nonLog := d.nonLoggedHours()
	nonLog(2)
	nonLog(3)
	nonLog(5)

	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	d.LogHours(Wednesday, 10)
	d.LogHours(Thursday, 10)
	d.LogHours(Friday, 6)
	d.LogHours(Saturday, 8)

	d.PayDetails()
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

func (d *Developer) nonLoggedHours() func(int) int {
	totalNonLogged := 0
	return func(hour int) int {
		return totalNonLogged + hour
	}
}

func (d *Developer) PayDay() (int, bool) {
	weekHour := d.HoursWorked()
	if weekHour <= 40 {
		return weekHour * d.HourlyRate, false
	} else {
		return 40*d.HourlyRate + (weekHour-40)*2*d.HourlyRate, true
	}
}

func (d *Developer) PayDetails() {
	for i := Monday; i <= Sunday; i++ {
		fmt.Printf("%v hours: %v\n", weekday[i], d.WorkWeek[i])
	}
	fmt.Println("Hours worked this week: ", d.HoursWorked())
	payday, overtime := d.PayDay()
	fmt.Println("Pay for the week: $", payday)
	fmt.Println("Is this overtime pay: ", overtime)
}
