package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	s "user/pkg/service"
)

func main() {
	mux := http.NewServeMux()
	srv := s.NewService()

	mux.HandleFunc("/create", srv.Create)                   // create a new user
	mux.HandleFunc("/make/friends", srv.MakeFriends)        // make two users frineds
	mux.HandleFunc("/delete/user{id}", srv.DeleteUser)      // delete user by id
	mux.HandleFunc("/get", srv.GetAll)                      // get all users
	mux.HandleFunc("/get/user{id}/friends", srv.GetFriends) // get friends of the user
	mux.HandleFunc("/get/user{id}/putAge", srv.ChangeAge)   // change the age of the user

	http.ListenAndServe("localhost:8080", mux)
}

func userWork() {
	fmt.Println("What do you wanna to do? (Choose number):")
	fmt.Println("1. Create new user")
	fmt.Println("2. Make two users friends")
	fmt.Println("3. Delete user")
	fmt.Println("4. Get all friends of the user")
	fmt.Println("5. Change the age of the user")
	fmt.Println("6. Exit")

	for {
		answer := ""
		fmt.Scanf("%s", &answer)
		if answer == "1" {
			fmt.Println("Give me the name, age and array of friends:")
			fmt.Scanf("%s", &answer)
			content := strings.Split(answer, " ")
			if len(content) < 3 {
				fmt.Println("More fields needed")
				continue
			}
			name := content[0]
			age := content[1]
			friends := content[2]
			fmt.Println(
				exec.Command(
					"curl", "-X POST -d '{\"name\":", name,
					"age\":", age, "friends\":", friends, "}' \"localhost:8080/create\"").
					Output())
		}
	}
}
