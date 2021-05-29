package main

import (
	"fmt"
	"runtime"
)

func main() {
	var list []int
	//var list []int
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)

	var a int = 100
	var b float32 = 100
	var c float64 = 100
	fmt.Println((a / 3) * 3)
	fmt.Println((b / 3) * 3)
	fmt.Println((c / 3) * 3)
}
