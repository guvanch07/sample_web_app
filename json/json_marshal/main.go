package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func main() {
	user1 := User{Id: 12, Name: "John"}
	bytes, _ := json.Marshal(user1)
	fmt.Println(string(bytes))
}
