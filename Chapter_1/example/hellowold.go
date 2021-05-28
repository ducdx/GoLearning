package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	Debug       bool      = false
	LogLevel    string    = "debug"
	startUpTime time.Time = time.Now()
)

func main() {
	debug, loglevel, starttime := false, "info", time.Now()

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	star := strings.Repeat("*", r)
	fmt.Println(star)

	x, y, z := getConfig()
	fmt.Println(x, y, z)

	fmt.Println(debug, loglevel, starttime)
}

func getConfig() (bool, string, time.Time) {
	return Debug, LogLevel, startUpTime
}
