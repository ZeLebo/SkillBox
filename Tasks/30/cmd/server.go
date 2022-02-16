package main

import (
	"net/http"
	s "user/pkg/service"
)

func main() {
	mux := http.NewServeMux()
	srv := s.NewService()

	mux.HandleFunc("/getUsers", srv.GetAll)                 // get all users
	mux.HandleFunc("/create", srv.Create)                   // create a new user
	mux.HandleFunc("/get/user{id}/putAge", srv.ChangeAge)   // change the age of the user
	mux.HandleFunc("/get/user{id}/friends", srv.GetFriends) // get friends of the user
	mux.HandleFunc("/make/friends", srv.MakeFriends)        // make two users friends
	mux.HandleFunc("/delete/user{id}", srv.DeleteUser)      // delete user by id

	http.ListenAndServe("localhost:8080", mux)
}
