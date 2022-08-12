package server

import (
	"github.com/gorilla/mux"
	"net/http"
	service "user/pkg/service"
)

// MyHandler defines the routes, returns router
func MyHandler() *mux.Router {
	srv := service.NewService()
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			return
		}
	}).Methods("GET")

	// request handlers
	router.HandleFunc("/get", srv.GetAllUsers).Methods("GET")                             // get all users
	router.HandleFunc("/create", loggerMiddleware(srv.Create)).Methods("POST")            // create a new user
	router.HandleFunc("/make_friends", loggerMiddleware(srv.MakeFriends)).Methods("POST") // make two users friends
	router.HandleFunc("/user", loggerMiddleware(srv.DeleteUser)).Methods("DELETE")        // delete user by target_id

	router.HandleFunc("/{id:[0-9]+}", loggerMiddleware(srv.ChangeAge)).Methods("PUT") // change the age of the user
	router.HandleFunc("/friends/{id:[0-9]+}", srv.GetFriends).Methods("GET")          // get friends of the user
	return router
}
