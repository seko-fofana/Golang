package main

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
)

func main() {
	type User struct {
		Login    string `json:"userName"`
		Password string
	}

	fich, err := ioutil.ReadFile("user.json")
	if err != nil {
		fmt.Println("error:", err)
	}
	var users []User
	err = json.Unmarshal(fich, &users)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v", users)
}
