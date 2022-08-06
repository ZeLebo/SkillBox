package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	s "user/pkg/service"
)

// The entry point

func main() {
	srv := s.NewService()
	router := mux.NewRouter()

	// request handlers
	router.HandleFunc("/get", srv.GetAllUsers)          // get all users
	router.HandleFunc("/create", srv.Create)            // create a new user
	router.HandleFunc("/make_friends", srv.MakeFriends) // make two users friends
	router.HandleFunc("/user", srv.DeleteUser)          // delete user by target_id

	router.HandleFunc("/{id:[0-9]+}", srv.ChangeAge)          // change the age of the user
	router.HandleFunc("/friends/{id:[0-9]+}", srv.GetFriends) // get friends of the user

	log.Error(http.ListenAndServe("localhost:8080", nil))
}
