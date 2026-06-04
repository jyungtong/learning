package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{Name: "John", Age: 21},
		{Name: "Lily", Age: 30},
	}

	marUsers, err := json.Marshal(users)
	if err != nil {
		log.Fatalln(err)
		return
	}

	os.WriteFile("users.json", marUsers, 0644)

	usersFileByte, err := os.ReadFile("users.json")
	if err != nil {
		log.Fatalln(err)
		return
	}

	var readUsers []User
	err = json.Unmarshal(usersFileByte, &readUsers)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(readUsers)
}
