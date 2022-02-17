// Package user /*
package user

import (
	"fmt"
	"reflect"
	"strconv"
)

type User struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends []*User `json:"friends"`
}

func (u *User) ToString() string {
	result := fmt.Sprintf("Name is %s, Age is %d", u.Name, u.Age)
	if len(u.Friends) > 0 {
		if len(u.Friends) > 1 {
			result += ", Friends ["
			for i, man := range u.Friends {
				result += "{"
				result += man.GetName() + " "
				result += strconv.Itoa(man.GetAge()) + "}"
				if i != len(u.Friends)-1 {
					result += ", "
				}
			}
			result += "]\n"
		} else {
			result += " Friend {"
			result += u.Friends[0].GetName() + " "
			result += strconv.Itoa(u.Friends[0].GetAge()) + "}\n"
		}
	}
	return result
}

// GetName returns the name of the user
func (u User) GetName() string {
	return u.Name
}

// GetAge returns the age of the user
func (u User) GetAge() int {
	return u.Age
}

// GetFriends returns the friends of the user
func (u User) GetFriends() []*User {
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

func (u *User) isFriend(friend *User) bool {
	for _, i := range u.Friends {
		if (i.Name == friend.Name) && (i.Age == friend.Age) {
			for j := range i.Friends {
				if reflect.DeepEqual(friend, i.Friends[j]) {
					return false
				}
			}
			return true
		}
	}

	return false
}

func (u *User) AddFriend(friend *User) bool {
	if u.isFriend(friend) {
		return false
	}
	u.Friends = append(u.Friends, friend)
	return true
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
}

// RemoveFriends remove one or more friends from user
func (u *User) RemoveFriends(friends ...User) {
	for i, man := range friends {
		for j := range u.Friends {
			if reflect.DeepEqual(u.Friends[j], man) {
				u.Friends = append(u.Friends[:i], u.Friends[i+1:]...)
				break
			}
		}
	}
}

// NewUser creates a new user and returns the link to it
func NewUser(name string, age int) User {
	user := User{
		Name: name,
		Age:  age,
	}
	return user
}

// FriendPeople function to make two users friends
func FriendPeople(u1, u2 *User) {
	u1.Friends = append(u1.Friends, u2)
	u2.Friends = append(u2.Friends, u1)
}
