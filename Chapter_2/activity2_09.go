package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		number := rand.Intn(8)
		if number%3 == 0 {
			fmt.Println("Skipping")
			continue
		} else if number%2 == 0 {
			fmt.Println("Stopping")
			break
		}

		fmt.Println(number)
	}

}
