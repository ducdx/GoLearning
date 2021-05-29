package main

import "fmt"

func main() {
	var a int8 = 125
	for i := 0; i < 20; i++ {
		a += 1
		fmt.Println(a)
	}
}
