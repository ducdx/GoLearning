package main

import (
	"errors"
	"fmt"
)

func main() {
	errBadData := errors.New("bad data")
	fmt.Printf("ErrBadData type: %T -- value: %v", errBadData, errBadData)
}
