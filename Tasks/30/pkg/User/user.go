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

func (u User) GetName() string {
	return u.Name
}

func (u User) GetAge() int {
	return u.Age
}

func (u User) GetFriends() []User {
	return u.Friends
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func (u *User) AddFriends(friends ...User) {
	if u.Friends == nil {
		u.Friends = make([]User, 0)
	}
	u.Friends = append(u.Friends, friends...)
}

func (u *User) ClearFriends() {
	u.Friends = nil
}

func (u *User) RemoveFriend(friend User) {
	for i := range u.Friends {
		if reflect.DeepEqual(u.Friends[i], friend) {
			u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
			return
		}
	}
	log.Printf("No such friend")
}

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

func (u *User) EraseUser() {
	log.Println("Removing", u.Name)
	u.Friends = nil
	u = nil
}

func NewUser(name string, age int) *User {
	user := &User{Name: name, Age: age}
	log.Println("New user", user.Name, "was make")
	return user
}
