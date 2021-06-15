package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(whatstheclock())

	TimeToManupulate := time.Now()
	ToBeAdd := time.Duration(10 * time.Minute)
	fmt.Println("The original time: ", TimeToManupulate)
	fmt.Println(ToBeAdd, " duration later: ", TimeToManupulate.Add(ToBeAdd))

	Start := time.Now()
	time.Sleep(2 * time.Second)
	End := time.Now()
	fmt.Println(elapsedTime(Start, End))

	date := time.Date(2021, 6, 15, 16, 30, 0, 0, time.Now().Local().Location())
	fmt.Println(date)
	next := date.AddDate(1, 1, 2)
	fmt.Println(next)
}

func whatstheclock() string {
	return time.Now().Format(time.ANSIC)
}

func elapsedTime(start time.Time, end time.Time) string {
	Elapsed := end.Sub(start)
	Hours := strconv.Itoa(int(Elapsed.Hours()))
	Minutes := strconv.Itoa(int(Elapsed.Minutes()))
	Seconds := strconv.Itoa(int(Elapsed.Seconds()))
	return "The total execution time elapsed is: " + Hours + " hour(s) and " + Minutes + " minute(s) and " + Seconds + " second(s)"
}
