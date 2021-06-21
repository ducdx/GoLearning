package main

import "fmt"

var budgetCatgories = make(map[int]string)

func init() {
	fmt.Println("Initializing our budgetCatgories...")
	budgetCatgories[1] = "Car Insurance"
	budgetCatgories[2] = "Mortgage"
	budgetCatgories[3] = "Electricity"
	budgetCatgories[4] = "Retirement"
	budgetCatgories[5] = "Vacation"
	budgetCatgories[6] = "Groceries"
	budgetCatgories[7] = "Car Payment"
}

func main() {
	for k, v := range budgetCatgories {
		fmt.Printf("key: %d -- Value: %s\n", k, v)
	}
}
