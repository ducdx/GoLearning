package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hour   int    `json:"courshours"`
}

type students struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"minitial"`
	FirstName     string   `json:"fname"`
	IsEnrolled    bool     `json:"enrolled"`
	Courses       []course `json:"classes"`
}

func main() {
	data := []byte(`
    {
      "id": 123,
      "lname": "Smith",
      "minitial": null,
      "fname": "John",
      "enrolled": true,
      "classes": [{
        "coursename": "Intro to Golang",
        "coursenum": 101,
        "coursehours": 4
      },
    {
        "coursename": "English Lit",
        "coursenum": 101,
        "coursehours": 3
      },
    {
        "coursename": "World History",
        "coursenum": 101,
        "coursehours": 3
      }
  ]
    }
  `)

	//Decode json
	var s students

	if !json.Valid(data) {
		fmt.Println("Json data is invalid")
		os.Exit(1)
	}

	err := json.Unmarshal(data, &s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}

	//Encode Json
	classes := []course{{
		Name:   "Go Programming",
		Number: 20,
		Hour:   24,
	}, {
		Name:   "Python Programming",
		Number: 25,
		Hour:   34,
	}}

	student := students{
		StudentId:     1111,
		LastName:      "Do",
		MiddleInitial: "Xuan",
		FirstName:     "Duc",
		IsEnrolled:    true,
		Courses:       classes,
	}

	js, err2 := json.MarshalIndent(student, "", "    ")

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("%s", js)
	}
}
