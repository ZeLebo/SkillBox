package main

import (
	"net/http"
	s "user/pkg/service"
)

func main() {
	mux := http.NewServeMux()
	srv := s.NewService()

	mux.HandleFunc("/get", srv.GetAllUsers)          // get all users
	mux.HandleFunc("/create", srv.Create)            // create a new user
	mux.HandleFunc("/make_friends", srv.MakeFriends) // make two users friends
	mux.HandleFunc("/user", srv.DeleteUser)          // delete user by target_id

	// The request has to be done this way
	// curl -X PUT -d '{"new age":1000}' "localhost:8080?id=1298498081"
	mux.HandleFunc("/", srv.ChangeAge)          // change the age of the user
	mux.HandleFunc("/friends/", srv.GetFriends) // get friends of the user

	http.ListenAndServe("localhost:8080", mux)
}
