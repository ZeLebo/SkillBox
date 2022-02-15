// Package user /*
package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	u "user/pkg/user"
)

type service struct {
	store map[int32]*u.User
}

func NewService() *service {
	return &service{make(map[int32] *u.User)}
}

func (s *service) EraseUser(u *u.User) {
	var tmp int32
	for id, user := range s.store {
		if user == u {
			tmp = id
		}
	}
	delete(s.store, tmp)
}

func (s *service) newId() int32 {
	var id int32
	for s.store[id] != nil {
		id = rand.Int31() % 100000
	}
	return id
}

// Create function to create a user
// return id of user
func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err = w.Write([]byte(err.Error())); err != nil {
				log.Println("Cannot write internal error")
			}
		}

		tmpUser := u.NewUser("", 0)

		// What if friends are new users?
		if err := json.Unmarshal(content, &tmpUser); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if _, err = w.Write([]byte(err.Error())); err != nil {
				log.Println("Cannot write internal error")
			}
		} else {
			fmt.Println(string(content))
		}

		id := s.newId()
		s.store[id] = tmpUser

		w.WriteHeader(http.StatusCreated)
		if _, err = w.Write([]byte("\nUser " + tmpUser.Name + " was created\nid:" +
			fmt.Sprintf("%x", id) + "\n")); err != nil {
			log.Println("Cannot write created user")
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

// MakeFriends make friends from 2 users
func (s *service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte(err.Error())); err != nil {

		}
		return
	}
	// debug session
	w.Write(content)

	if _, err = w.Write([]byte("Users are now friends\n")); err != nil {
		log.Println("Something went wrong")
	}
}

func (s *service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// TODO how to delete specified user?

	w.Write(content)
	if _, err := w.Write([]byte("User has been deleted\n")); err != nil {
		log.Println("User has been deleted")
	}
}

func (s *service) GetFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s *service) ChangeAge(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s *service) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response := ""
	for _, user := range s.store {
		response += user.ToString()
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(response)); err != nil {
		log.Println()
	}
	return
}
