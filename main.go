package main

import (
	"encoding/json"
	"flag"
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

func (users *Users) AddUser(user User) []User  {
	users.Users = append(users.Users, user)
	return users.Users
}

func (users *Users) deleteUser(id int) []User  {
	for i, user := range users.Users {
		if user.Id == id {
			users.Users = append(users.Users[:i], users.Users[i+1:]...)
			break
		}
	}
	return users.Users
}

func main() {
	types := flag.String("type", "empty", "available type: create | read | update | delete")
	file, err := os.ReadFile("./users.json")
	flag.Parse()	

	if err != nil {
		log.Fatal(err)
		return
	}

	var users Users
	var username string
	var email string
	var password string
	
	json.Unmarshal(file, &users)

	if *types != "empty" {
		switch *types {
		case "create":
			var newUser User
			fmt.Print("Username: ")
			fmt.Scanln(&username)
			fmt.Print("Email: ")
			fmt.Scanln(&email)
			fmt.Print("Password: ")
			fmt.Scanln(&password)
			newUser.Id = users.Users[len(users.Users)-1].Id + 1
			newUser.Username = username
			newUser.Email = email
			newUser.Password = password
			users.AddUser(newUser)
			appendUser,_ := json.Marshal(users)
			err := os.WriteFile("./users.json", []byte(appendUser), 0666)
			if err != nil {
				log.Fatal("Some Error Occured!")
				return
			}
			break
		case "readByID":
			break
		case "readAll":
			for _, user := range users.Users {
				println("ID: ", user.Id)
				println("Username: ", user.Username)
				println("Email: ", user.Email)
				println("Password: ", user.Password, "\n")
			}
			break
		case "update":
			break
		case "delete":
			var id int
			fmt.Print("ID: ")
			fmt.Scanln(&id)
			users.deleteUser(id)
			deleteUser,_ := json.Marshal(users)
			err := os.WriteFile("./users.json", []byte(deleteUser), 0666)
			if err != nil {
				log.Fatal("Some Error Occured!")
				return
			}
			break
		default:
			log.Fatal("Error: Option are not available")
			return
		}
	}
}