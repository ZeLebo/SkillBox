package database

import (
	"gorm.io/gorm"
	u "user/internal/user"
)

type IClient interface {
	Connect() (*gorm.DB, error)
	// Connect() (*sql.DB, error)
	GetUsers() ([]u.User, error)
	CreateUser(Model) (int, error)
	DeleteUser(int) error
	ChangeAge(int, int) error
	MakeFriends(int, int) error
	GetFriends(int) ([]u.User, error)
	CheckUser(Model) bool
	GetUserByID(int) (u.User, error)
	GetUserID(Model) int
}
