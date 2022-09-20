// Package Service /*
package service

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
	"user/internal/database"
	"user/internal/logger"
	"user/internal/service/validators"
	u "user/internal/user"
)

type IService interface {
	Contains(*u.User) bool
	GetAllUsers(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	ChangeAge(http.ResponseWriter, *http.Request)
	GetFriends(http.ResponseWriter, *http.Request)
	MakeFriends(http.ResponseWriter, *http.Request)
}

// TODO postgresql database instead of map
type Service struct {
	store map[int32]*u.User
}

func NewService(client *database.Client) *Service {
	return &Service{}
}

// Contains check if the map Contains the specific user
func (s *Service) Contains(u *u.User) bool {
	for _, i := range s.store {
		if i == u {
			return true
		}
	}
	return false
}

// Id generator
func (s *Service) newId() int32 {
	var id int32
	// It's limited to 2^31 + 1
	// Wanted to use hash, but then thought it would be too much
	for s.store[id+1] != nil {
		id = rand.Int31()
	}
	return id + 1
}

// GetAllUsers func to return all the users in the map
func (s *Service) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	// collecting data from the database
	// here's request to the database that returns list of users
	for id, user := range s.store {
		_, err := w.Write([]byte("id: " + strconv.Itoa(int(id)) +
			"\nuser: " + user.ToString() + "\n"))
		if err != nil {
			return
		}
	}
}

// Create function returns id of user
func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	var req validators.Request
	err := req.Bind(r)
	if err != nil {
		logger.HTTPErrorHandle(w, logger.HTTPErrorHandler{
			ErrorCode:   http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	tmpUser := u.NewUser(req.Name, req.Age)

	id := s.newId()
	s.store[id] = &tmpUser

	// What if friends are new users? -> make new users
	func(u []*u.User) {
		for _, man := range u {
			if !s.Contains(man) {
				newId := s.newId()
				s.store[newId] = man
			}
		}
	}(s.store[id].GetFriends())

	log.Info("New user: ", id)
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("\nuser " + tmpUser.Name + " was created\nid:" + strconv.Itoa(int(id)) + "\n"))
	if err != nil {
		return
	}
}

// ChangeAge to change the age of specific user
func (s *Service) ChangeAge(w http.ResponseWriter, r *http.Request) {
	var req validators.Request
	err := req.Bind(r)
	if err != nil {
		logger.HTTPErrorHandle(w, logger.HTTPErrorHandler{
			ErrorCode:   http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	// parse the header of request
	vars := mux.Vars(r)
	tmp, _ := strconv.Atoi(vars["id"])
	req.TargetID = int32(tmp)

	if _, ok := s.store[req.TargetID]; !ok {
		_, err := w.Write([]byte("No such user"))
		if err != nil {
			return
		}
		return
	}
	// change age doesn't change the age in friends
	s.store[req.TargetID].SetAge(req.Age)
	_, err = w.Write([]byte("user's age was updated\n"))
	if err != nil {
		return
	}
	log.Info("user ", req.TargetID, " age has been changed to ", req.Age)
}

// GetFriends of specific user
func (s *Service) GetFriends(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmp, _ := strconv.Atoi(vars["id"])
	id := int32(tmp)

	if _, ok := s.store[id]; !ok {
		_, err := w.Write([]byte("No such user"))
		if err != nil {
			return
		}
		return
	}
	// collecting data from the user
	answer := func(u []*u.User) string {
		if len(u) == 0 {
			return "user has no friends\n"
		}
		result := "Friends of " + s.store[id].Name + ":"
		for _, man := range u {
			result += "\n" + man.ToString()
		}
		return result + "\n"
	}(s.store[id].GetFriends())

	_, err := w.Write([]byte(answer))
	if err != nil {
		return
	}
}

// MakeFriends make friends from 2 users
func (s *Service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	var data validators.Request
	err := data.Bind(r)
	if err != nil {
		logger.HTTPErrorHandle(w, logger.HTTPErrorHandler{
			ErrorCode:   http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}

	// id cannot be < 1, so if we have 0 it means user hasn't provided us the fields
	if data.TargetID == 0 || data.SourceID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("You need to provide the id's of users"))
		if err != nil {
			return
		}
		return
	}

	if data.TargetID == data.SourceID {
		_, err := w.Write([]byte("The same user"))
		if err != nil {
			return
		}
		return
	}

	// We need to store the previous copy of the user
	// to exclude the collision
	var tmp u.User
	tmp.Name = s.store[data.TargetID].GetName()
	tmp.Age = s.store[data.TargetID].GetAge()
	tmp.Friends = s.store[data.TargetID].GetFriends()

	// AddFriends returns true if succeed to add a new user
	// so if true, we can add another user to friends list
	if s.store[data.TargetID].AddFriend(s.store[data.SourceID]) {
		s.store[data.SourceID].AddFriend(&tmp)
	} else {
		// if false -> we already have such user in the map
		_, err := w.Write([]byte("Users are already friends\n"))
		if err != nil {
			return
		}
		return
	}

	// just printing and logging
	_, err = w.Write([]byte("Users " + s.store[data.TargetID].GetName() + " and " +
		s.store[data.SourceID].GetName() + " are now friends\n"))
	if err != nil {
		return
	}
	log.Info("Users ", s.store[data.TargetID].GetName()+" and "+
		s.store[data.SourceID].GetName(), " are now friends")
}

// DeleteUser from the map
func (s *Service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var data validators.Request
	err := data.Bind(r)
	if err != nil {
		logger.HTTPErrorHandle(w, logger.HTTPErrorHandler{
			ErrorCode:   http.StatusBadRequest,
			Description: err.Error(),
		})
		return
	}
	// check if this user exists
	if _, ok := s.store[data.TargetID]; !ok {
		_, err := w.Write([]byte("No such user"))
		if err != nil {
			return
		}
		return
	}

	// delete from all the friends
	for _, man := range s.store[data.TargetID].GetFriends() {
		man.RemoveFriend(*s.store[data.TargetID])
	}
	// logging and deleting
	_, err = w.Write([]byte("user " + s.store[data.TargetID].GetName() + " has been deleted\n"))
	if err != nil {
		return
	}
	log.Info("user " + s.store[data.TargetID].GetName() + " has been deleted")
	delete(s.store, data.TargetID)
}
