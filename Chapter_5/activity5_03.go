package main

import "fmt"

func main() {
	i := 0

	incrementor := func() int {
		i += 1
		return i
	}

	fmt.Println(incrementor())
	fmt.Println(incrementor())

	i += 10
	fmt.Println(incrementor())
}
