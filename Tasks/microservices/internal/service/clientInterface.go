package service

import (
	"gorm.io/gorm"
	u "user/internal/user"
)

type ClientInterface interface {
	Connect() (*gorm.DB, error)
	GetUsers() ([]u.User, error)
	CreateUser(u.User) (int, error)
	DeleteUser(int) error
	ChangeAge(int, int) error
	MakeFriends(int, int) error
	GetFriends(int) ([]u.User, error)
	CheckUser(u.User) bool
	GetUserByID(int) (u.User, error)
	GetUserID(u.User) int
}
