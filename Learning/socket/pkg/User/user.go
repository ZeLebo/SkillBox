/*
This is the user structure
*/
package user

import (
	"log"
	"reflect"
) 

type User struct {
	Name    string
	Age     int
	Friends []User
}

// returns the name of the user
func (u User) GetName() string {
	return u.Name
}

// returns the age of the user
func (u User) GetAge() int {
	return u.Age
}

// returns the friends of the user
func (u User) GetFriends() []User {
	return u.Friends
}

// set the name of the user
func (u *User) SetName(name string) {
	u.Name = name
}

// set the age of the user
func (u *User) SetAge(age int) {
	u.Age = age
}

// add one or more friends to the user
func (u *User) AddFriends(friends ...User) {
	if u.Friends == nil {
		u.Friends = make([]User, 1)
	}
	u.Friends = append(u.Friends, friends...)
}

// remove all the friends from the user
func (u *User) ClearFriends() {
	u.Friends = nil
}

// remove one friend from user
func (u *User) RemoveFriend(friend User) {
	for i := range u.Friends {
		if reflect.DeepEqual(u.Friends[i], friend) {
			u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
			return
		}
	}
	log.Printf("No such friend")
}

// remove one or more friends from user
func (u *User) RemoveFriends(friends ...User) {
	for i := range friends {
		for j := range u.Friends {
			if reflect.DeepEqual(u.Friends[j], i) {
				u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
				break
			}
		}
	}
}

// erases the user from the memory
func (u *User) EraseUser() {
	log.Println("Removing", u.Name)
	u.Friends = nil
	u = nil
}

// creates a new user and returns the link to it
func NewUser(name string, age int) *User {
	user := &User{Name: name, Age: age}
	log.Println("New user", user.Name, "was make")
	return user
}
