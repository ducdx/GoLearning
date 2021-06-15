package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

type directDeposit struct {
	firstName     string
	lastName      string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validateLastName() error {
	if d.lastName == "" {
		return ErrInvalidLastName
	}
	return nil
}

func (d *directDeposit) validateRoutingNumber() error {
	if d.routingNumber < 100 {
		return ErrInvalidRoutingNumber
	}
	return nil
}

func (d *directDeposit) report() {
	err := d.validateLastName()
	if err != nil {
		fmt.Println(err)
	}
	err = d.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("***********************************************************")
	fmt.Println("First Name: ", d.firstName)
	fmt.Println("Last Name: ", d.lastName)
	fmt.Println("Bank Name: ", d.bankName)
	fmt.Println("Routing Number:", d.routingNumber)
	fmt.Println("Account Number: ", d.accountNumber)
}

func main() {
	d := directDeposit{
		firstName:     "Duc",
		lastName:      "",
		bankName:      "Vietcombank",
		routingNumber: 17,
		accountNumber: 1234567,
	}
	d.report()
}
