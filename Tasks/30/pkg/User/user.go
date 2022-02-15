/*
This is the user structure
*/

package user

import (
	"fmt"
	"log"
	"reflect"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []User `json:"friends"`
}

func (u *User) ToString() string {
	return fmt.Sprintf("Name is %s, Age is %d, Friends: %v\n", u.Name, u.Age, u.Friends)
}

// GetName returns the name of the user
func (u User) GetName() string {
	return u.Name
}

// GetAge returns the a™ge of the user
func (u User) GetAge() int {
	return u.Age
}

// GetFriends returns the friends of the user
func (u User) GetFriends() []User {
	return u.Friends
}

// SetName set the name of the user
func (u *User) SetName(name string) {
	u.Name = name
}

// SetAge set the age of the user
func (u *User) SetAge(age int) {
	u.Age = age
}

// AddFriends add one or more friends to the user
func (u *User) AddFriends(friends ...User) {
	if u.Friends == nil {
		u.Friends = make([]User, 1)
	}
	u.Friends = append(u.Friends, friends...)
}

// ClearFriends remove all the friends from the user
func (u *User) ClearFriends() {
	u.Friends = nil
}

// RemoveFriend remove one friend from user
func (u *User) RemoveFriend(friend User) {
	for i := range u.Friends {
		if reflect.DeepEqual(u.Friends[i], friend) {
			u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
			return
		}
	}
	log.Printf("No such friend")
}

// RemoveFriends remove one or more friends from user
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

// NewUser creates a new user and returns the link to it
func NewUser(name string, age int) *User {
	user := &User{Name: name, Age: age}
	log.Println("New user was created")
	return user
}

// FriendPeople function to make two users friends
func FriendPeople(u1, u2 *User) {
	u1.Friends = append(u1.Friends, *u2)
	u2.Friends = append(u2.Friends, *u1)
}
