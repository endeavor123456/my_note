package main

import (
	"encoding/json"
	"fmt"
)

type MessageInfo struct {
}
type Name struct {
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
}
type People struct {
	Name    Name     `json:"name"`
	Age     int      `json:"age"`
	Weight  float64  `json:"weight"`
	Message []string `json:"message"`
}

func main() {
	var a1 = []People{ //a1是切片类型的变量
		{
			Name: Name{
				Lastname:  "李",
				Firstname: "二",
			},
			Age:    20,
			Weight: 70.4,
			Message: []string{
				"Good morning.", "How are you?", "Nice to meet you."},
		},
		{
			Name: Name{
				Lastname:  "王",
				Firstname: "五",
			},
			Age:    33,
			Weight: 80.3,
			Message: []string{
				"Good morning.", "How are you?", "Nice to meet you."},
		},
	}

	data, err := json.MarshalIndent(a1, "", "\t")
	if err == nil {
		fmt.Printf("%s\n", data)
	}
}
