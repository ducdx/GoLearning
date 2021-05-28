package main

import "fmt"

func main() {
	wordMap := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}
	mostPopular := 0
	mostWord := ""
	for key, val := range wordMap {
		if mostPopular < val {
			mostPopular = val
			mostWord = key
		}
	}
	fmt.Println("Most popular word is: ", mostWord, " with count: ", mostPopular)
}
