package main

import "fmt"

func main() {
	counter := 4
	decre := decrementor(counter)
	fmt.Println(decre())
	fmt.Println(decre())
	fmt.Println(decre())
}

func decrementor(i int) func() int {
	return func() int {
		i -= 1
		return i
	}
}
