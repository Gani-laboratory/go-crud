package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Id int
	Username string
	Email string
	Password string
}
type Users struct {
	Users []User
}

func main() {
	var users Users
	file, err := os.ReadFile("./users.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(file, &users)
	fmt.Println(users.Users)
}