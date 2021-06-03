package main

import "fmt"

func main() {
	calculator(add, 5, 6)
	calculator(subtract, 10, 4)
}

func add(i, j int) int {
	return i + j
}

func subtract(i, j int) int {
	return i - j
}

func calculator(f func(i, j int) int, i, j int) {
	fmt.Println(f(i, j))
}
