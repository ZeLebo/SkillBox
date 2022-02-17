// Package service /*
package service

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	u "user/pkg/user"
)

// There are 20 warnings for "write" function
// It will be better to check error while writing
// But I don't wanna do this

type request struct {
	TargetID int32 `json:"target_id"`
	SourceID int32 `json:"source_id"`
	Age      int   `json:"new age"`
}

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

func (s *service) newId() int32 {
	var id int32
	for s.store[id+1] != nil {
		id = rand.Int31()
	}
	return id + 1
}

func (s *service) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for id, user := range s.store {
		w.Write([]byte("id: " + strconv.Itoa(int(id)) + "\nUser: " + user.ToString() + "\n"))
	}
	//w.WriteHeader(http.StatusOK) // To print to console
}

// Create function returns id of user
func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	content, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Cannot read the data from request")
		w.Write([]byte(err.Error()))
	}

	tmpUser := u.NewUser("", 0)

	if err := json.Unmarshal(content, &tmpUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Error("Cannot parse data from json")
	}

	id := s.newId()
	s.store[id] = &tmpUser

	// What if friends are new users?
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
	w.Write([]byte("\nUser " + tmpUser.Name + " was created\nid:" + strconv.Itoa(int(id)) + "\n"))
}

func (s *service) ChangeAge(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req request
	content, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Error("Wrong request to change age")
		return

	}
	if err := json.Unmarshal(content, &req); err != nil {
		log.Error("Cannot parse request to change age")
		return
	}

	tmp := r.URL.Query().Get("id")
	num, _ := strconv.Atoi(tmp)
	req.TargetID = int32(num)

	if _, ok := s.store[req.TargetID]; !ok {
		w.Write([]byte("No suck user"))
		return
	}
	// change age doesn't change the age in friends
	s.store[req.TargetID].SetAge(req.Age)
	w.Write([]byte("User's age was updated"))
	log.Info("User ", req.TargetID, " age has been changed to ", req.Age)
}

func (s *service) GetFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tmp := r.URL.Query().Get("id")
	num, _ := strconv.Atoi(tmp)
	id := int32(num)

	if _, ok := s.store[id]; !ok {
		w.Write([]byte("No such user"))
		return
	}
	answer := func(u []*u.User) string {
		result := "Friends of " + s.store[id].Name + ":"
		for _, man := range u {
			result += "\n" + man.ToString()
		}
		return result + "\n"
	}(s.store[id].GetFriends())

	w.Write([]byte(answer))
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
	var data request
	if err := json.Unmarshal(content, &data); err != nil {
		log.Error("Cannot parse data for making friends")
		w.Write([]byte(err.Error())) // or w.Write([]byte("Wrong request))
		return
	}

	if data.TargetID == 0 || data.SourceID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You need to provide the id's of users"))
	}

	if data.TargetID == data.SourceID {
		w.Write([]byte("The same user"))
		return
	}

	var tmp u.User
	tmp.Name = s.store[data.TargetID].GetName()
	tmp.Age = s.store[data.TargetID].GetAge()
	tmp.Friends = s.store[data.TargetID].GetFriends()

	// tmp := s.store[data.TargetID] // to exclude collision
	if s.store[data.TargetID].AddFriend(s.store[data.SourceID]) {
		s.store[data.SourceID].AddFriend(&tmp)
	} else {
		w.Write([]byte("Users are already friends\n"))
		return
	}

	w.Write([]byte("Users " + s.store[data.TargetID].GetName() + " and " +
		s.store[data.SourceID].GetName() + " are now friends\n"))
	log.Info("Users ", s.store[data.TargetID].GetName()+" and "+
		s.store[data.SourceID].GetName(), " are now friends")
}

func (s *service) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Wrong request for deleting user")
		return
	}
	var data request
	if err := json.Unmarshal(content, &data); err != nil {
		log.Error("Cannot parse data for deleting user")
		return
	}
	// check if this user exists
	if _, ok := s.store[data.TargetID]; !ok {
		w.Write([]byte("No such user"))
		return
	}

	// delete from all of the friends
	for _, man := range s.store[data.TargetID].GetFriends() {
		man.RemoveFriend(*s.store[data.TargetID])
	}
	w.Write([]byte("User " + s.store[data.TargetID].GetName() + " has been deleted\n"))
	log.Info("User " + s.store[data.TargetID].GetName() + " has been deleted")
	delete(s.store, data.TargetID)
}
