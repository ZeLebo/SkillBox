// Package service /*
package service

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	u "user/pkg/user"
)

type service struct {
	store map[int32]*u.User
}

func NewService() *service {
	return &service{make(map[int32]*u.User)}
}

func (s *service) Contains(u *u.User) bool {
	for _, i := range s.store {
		if i == u {
			return true
		}
	}
	return false
}

func (s *service) EraseUser(u *u.User) {
	for id, user := range s.store {
		if user == u {
			log.Info("User", u.Name, "has been erased")
			delete(s.store, id)
			return
		}
	}
}

func (s *service) newId() int32 {
	var id int32
	for s.store[id] != nil {
		id = rand.Int31() // TODO loop need to be fixed 2^31 + 1
	}
	return id
}

// Create function returns id of user
func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Cannot read the data from request")
		w.Write([]byte(err.Error()))
	}

	tmpUser := u.NewUser("", 0)

	// todo What if friends are new users?
	if err := json.Unmarshal(content, &tmpUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte(err.Error())); err != nil {
			log.Error("Cannot write internal error")
		}
	}

	id := s.newId()
	s.store[id] = &tmpUser

	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write([]byte("\nUser " + tmpUser.Name + " was created\nid:" +
		fmt.Sprintf("%x", id) + "\n")); err != nil {
		log.Info("Cannot write created user")
	}
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
		log.Error("Cannot parse request data")
		w.Write([]byte(err.Error()))
		return
	}
	// debug session
	w.Write(content)
	data := ""
	if err := json.Unmarshal(content, &data); err != nil {
		log.Error("Cannot friend ")
	}
	w.Write([]byte("Users are now friends\n"))
	log.Info("Users:")
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
