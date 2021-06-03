package main

import "fmt"

func main() {
	incre := incrementor()
	fmt.Println(incre())
	fmt.Println(incre())
	fmt.Println(incre())

	otherincre := incrementor()
	fmt.Println(otherincre())
}

func incrementor() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
