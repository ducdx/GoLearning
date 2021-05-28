package main

import "fmt"

func main() {
	arr := []int{1, 4, 6, 100, 20, 9, 7, 1000, 50, 68, 29, 54, 82, 15, 117, 32, 65, 38, 50, 98, 22, 38}
	fmt.Println("Before: ", arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("After: ", arr)
}
