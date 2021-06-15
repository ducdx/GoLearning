package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Before Panic")
	Test()
	fmt.Println("After Panic")
}

func Test() {
	n := func() {
		fmt.Println("Defer in test")
	}
	defer n()
	m := "good-bye"
	message(m)
}

func message(m string) {
	f := func() {
		fmt.Println("Defer in message func")
	}
	defer f()

	if m == "good-bye" {
		panic(errors.New("Something went wrong"))
	}
}
