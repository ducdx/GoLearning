package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	if count1 != nil {
		fmt.Printf("count 1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count 2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count 3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("time: %#v\n", t.String())
	}

	a, b := 5, 10
	fmt.Printf("Before: a=%d,b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After: a=%d,b=%d\n", a, b)
}

func swap(a *int, b *int) {
	*b, *a = *a, *b
}
