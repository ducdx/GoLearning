package main

import "fmt"

func main() {
	sl := []string{"Good", "Good", "Bad", "Good", "Good"}
	sl = append(sl[:2], sl[3:]...)
	fmt.Println("New Slice: ", sl)
}
