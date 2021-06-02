package main

import "fmt"

func main() {
	itemSold()
}

func itemSold() {
	sold := map[string]int{
		"John":  35,
		"Mike":  45,
		"Peter": 80,
		"Lan":   110,
	}

	for k, v := range sold {
		if v < 40 {
			fmt.Println(k, " sold", v, "items and below expectation")
		} else if v >= 40 && v < 100 {
			fmt.Println(k, " sold", v, " items and meets expectation")
		} else {
			fmt.Println(k, " sold", v, " items and exceed expectation")
		}
	}
}
