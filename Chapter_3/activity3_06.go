package main

import (
	"fmt"
)

func main() {
	name := "Tôi là Đỗ Xuân Đức"
	for i := 0; i < len(name); i++ {
		fmt.Print(name[i], " ")
	}
	fmt.Println()
	for i := 0; i < len(name); i++ {
		fmt.Print(string(name[i]), " ")
	}
	fmt.Println()
	runes := []rune(name)
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
	fmt.Println()

	for i, val := range runes {
		fmt.Println(i, string(val))
	}

	fmt.Println("Bytes: ", len(name))
	fmt.Println("Runes: ", len(runes))
	fmt.Println(string(name[:10]))
	fmt.Println(string(runes[:10]))
}
