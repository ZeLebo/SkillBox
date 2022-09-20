package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"user/configs"
	u "user/internal/user"
)

func NewPostgresDB(cfg *configs.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Client struct {
	db *sqlx.DB
}

func NewClient(cfg *configs.DBConfig) (*Client, error) {
	db, err := NewPostgresDB(cfg)
	if err != nil {
		return nil, err
	}
	return &Client{db: db}, nil
}

func (client *Client) Close() error {
	return client.db.Close()
}

func (client *Client) GetDB() *sqlx.DB {
	return client.db
}

func (client *Client) GetUsers() ([]u.User, error) {
	var users []u.User
	err := client.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (client *Client) CreateUser(user u.User) error {
	// add new user to the database with unique id, age and name
	_, err := client.db.Exec("INSERT INTO users (id, age, name) VALUES ($1, $2, $3)", user.ID, user.Age, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) DeleteUser(id int) error {
	_, err := client.db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) ChangeAge(id, age int) error {
	_, err := client.db.Exec("UPDATE users SET age=$1 WHERE id=$2", age, id)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) MakeFriends(id1, id2 int) error {
	_, err := client.db.Exec("INSERT INTO friends (id1, id2) VALUES ($1, $2)", id1, id2)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) GetFriends(id int) ([]u.User, error) {
	var users []u.User
	// maybe doesn't work
	err := client.db.Select(&users, "SELECT * FROM users WHERE id IN (SELECT id1 FROM friends WHERE id2=$1)", id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (client *Client) GetUserByID(id int) (u.User, error) {
	var user u.User
	err := client.db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return user, err
	}
	return user, nil
}
