package main

import "fmt"

type id string

func main() {
	v1 := 5
	fmt.Println(v1, " is type: ", getTypeValue(v1))
	v2 := 3.2323243
	fmt.Println(v2, " is type: ", getTypeValue(v2))
	v3 := "Demo"
	fmt.Println(v3, " is type: ", getTypeValue(v3))
	v4 := true
	fmt.Println(v4, " is type: ", getTypeValue(v4))
	v5 := struct {
		x int
		y float32
	}{
		10,
		5.12,
	}
	fmt.Println(v5, " is type: ", getTypeValue(v5))
}

func getTypeValue(val interface{}) string {
	str := ""
	switch val.(type) {
	case int, int32, int64:
		str = "int"
	case float32, float64:
		str = "float"
	case string:
		str = "string"
	case bool:
		str = "bool"
	default:
		str = "Unknown Type"
	}
	return str
}
