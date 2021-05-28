package main

import "fmt"

func main() {
	x := 5
	checkNumber(&x)
	x = 6
	checkNumber(&x)
}

func checkNumber(a *int) {
	if *a%2 == 0 {
		fmt.Println(*a, " is even!")
	} else {
		fmt.Println(*a, " is odd!")
	}
}
