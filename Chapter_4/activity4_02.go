package main

import (
	"fmt"
	"os"
)

func main() {
	m1 := map[string]string{
		"001": "Duc",
		"002": "QUynh",
		"003": "Nhim",
		"004": "Tho",
	}

	m2 := m1
	m1["003"] = "Hoai An"

	fmt.Println("Map 1: ", m1)
	fmt.Println("Map 2: ", m2)

	if len(os.Args) < 2 {
		fmt.Println("Argument not valid")
		os.Exit(1)
	}

	userId := os.Args[1]
	name, existed := m1[userId]
	if existed {
		fmt.Println(userId, " is Approved! Name is: ", name)
	} else {
		fmt.Println(userId, " is not Approved!")
	}

	for k, v := range m1 {
		fmt.Println("Key: ", k, " - Value: ", v)
	}
}
