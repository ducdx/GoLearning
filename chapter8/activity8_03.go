package main

import "fmt"

var budgetCatgories = make(map[int]string)
var payeeToCategory = make(map[string]int)

func init() {
	initBudgetCatgories()
	initPayeeToCategory()
}

func initBudgetCatgories() {
	fmt.Println("Initializing our budgetCatgories...")
	budgetCatgories[1] = "Car Insurance"
	budgetCatgories[2] = "Mortgage"
	budgetCatgories[3] = "Electricity"
	budgetCatgories[4] = "Retirement"
	budgetCatgories[5] = "Vacation"
	budgetCatgories[6] = "Groceries"
	budgetCatgories[7] = "Car Payment"
}

func initPayeeToCategory() {
	fmt.Println("Assign our Payees to categories")
	payeeToCategory["Nationwide"] = 1
	payeeToCategory["BBT Loan"] = 2
	payeeToCategory["First Energy Electric"] = 3
	payeeToCategory["Ameriprise Financial"] = 4
	payeeToCategory["Walt Disney World"] = 5
	payeeToCategory["ALDI"] = 7
	payeeToCategory["Martins"] = 7
	payeeToCategory["Wal Mart"] = 7
	payeeToCategory["Chevy Loan"] = 6
}

func main() {
	fmt.Println("In main, printing payee to category")
	for k, v := range payeeToCategory {
		fmt.Printf("Payee: %s, Category: %s\n", k, budgetCatgories[v])
	}
}
