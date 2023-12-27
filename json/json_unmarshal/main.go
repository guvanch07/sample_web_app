package main

import (
	"encoding/json"
	"fmt"
)

var DATA = `{
    "id": 44,
    "price": 4000,
    "items": [
        {
            "name": "snowBoard",
            "number": 1
        },
        {
            "name": "ski",
            "number": 3
        },
        {
            "name": "ball",
            "number": 3
        }
    ]
}`

type Order struct {
	Id    int    `json:"id"`
	Price int    `json:"price"`
	Items []Item `json:"items"`
	Zero  int
}
type Item struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func main() {
	var order1 Order
	err := json.Unmarshal([]byte(DATA), &order1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", order1)
}
