package main

import (
	"errors"
	"fmt"
)

func main() {
	a()
	fmt.Println("This line will now get printed from main() function")
}

func a() {
	b("good-bye")
	fmt.Println("Back in function a()")
}

func b(m string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in func b()", r)
		}
	}()

	if m == "good-bye" {
		panic(errors.New("Panic in func b()"))
	}
}
