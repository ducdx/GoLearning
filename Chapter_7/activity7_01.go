package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        int
	HoursWorkedInYear int
	Review            map[string]interface{}
}
type Manager struct {
	Individual    Employee
	Salary        int
	CommisionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (d *Developer) Pay() (string, float64) {
	name := d.Individual.FirstName + " " + d.Individual.LastName
	salary := d.HourlyRate * d.HoursWorkedInYear
	return name, float64(salary)
}

func (m *Manager) Pay() (string, float64) {
	name := m.Individual.FirstName + " " + m.Individual.LastName
	salary := float64(m.Salary) + float64(m.Salary)*m.CommisionRate
	return name, float64(salary)
}

func (d *Developer) calculateReview() float64 {
	if len(d.Review) == 0 {
		return 0
	}

	totalReview := 0
	for _, v := range d.Review {
		switch t := v.(type) {
		case int:
			totalReview += t
		case string:
			switch v {
			case "Excellent":
				totalReview += 5
			case "Good":
				totalReview += 4
			case "Fair":
				totalReview += 3
			case "Poor":
				totalReview += 2
			case "Unsatisfactor":
				totalReview += 1
			default:
				totalReview += 0
			}
		default:
			totalReview += 0
		}

	}
	return float64(totalReview) / float64(len(d.Review))
}

func (d *Developer) PayDetails() {
	name, pay := d.Pay()
	fmt.Println("total review:", len(d.Review))
	fmt.Printf("%s got a review rating of %.2f\n", name, d.calculateReview())
	fmt.Printf("%s got paid %v for the year \n\n", name, pay)
}

func (m *Manager) PayDetails() {
	name, pay := m.Pay()
	fmt.Printf("%s got paid %v for the year\n\n", name, pay)
}

func main() {
	d1 := Developer{
		Individual: Employee{
			FirstName: "Duc",
			LastName:  "Do",
		},
		HourlyRate:        50,
		HoursWorkedInYear: 2088,
		Review: map[string]interface{}{
			"1": 5,
			"2": "Excellent",
			"3": 4,
			"4": "Good",
			"5": 3,
			"6": "Fair",
			"7": "Fair",
		},
	}

	m1 := Manager{
		Individual: Employee{
			FirstName: "Manh",
			LastName:  "Nguyen",
		},
		Salary:        150000,
		CommisionRate: 1.5,
	}

	d1.PayDetails()
	m1.PayDetails()
}
