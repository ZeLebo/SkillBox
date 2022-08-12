package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"user/pkg/user"

	env "github.com/joho/godotenv"
)

type PGSQL struct {
	*sql.DB
}

func NewPGSQL(db *sql.DB) *PGSQL {
	userName := env.Load("POSTGRES_USER")
	//password := env.Load("POSTGRES_PASSWORD")
	dbName := env.Load("DB_NAME")
	connStr := fmt.Sprintf("user=%s dbname=%s", userName, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	return &PGSQL{db}
}

func (p *PGSQL) Close() {
	p.DB.Close()
}

func (p *PGSQL) GetUser(id int) (user user.User, err error) {
	err = p.QueryRow("SELECT id, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name)
	return
}

func (p *PGSQL) ChangeAge(id int, age int) (err error) {
	_, err = p.Exec("UPDATE users SET age = $1 WHERE id = $2", age, id)
	return
}

func (p *PGSQL) CreateUser(name string) (id int, err error) {
	err = p.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", name).Scan(&id)
	return
}

func (p *PGSQL) DeleteUser(id int) (err error) {
	_, err = p.Exec("DELETE FROM users WHERE id = $1", id)
	return
}

func (p *PGSQL) MakeFriends(id1 int, id2 int) (err error) {
	_, err = p.Exec("INSERT INTO friends (user_id, friend_id) VALUES ($1, $2)", id1, id2)
	return
}

func (p *PGSQL) GetFriends(id int) (friends []user.User, err error) {
	rows, err := p.Query("SELECT id, name FROM users WHERE id IN (SELECT friend_id FROM friends WHERE user_id = $1)", id)
	if err != nil {
		return
	}
	for rows.Next() {
		var friend user.User
		err = rows.Scan(&friend.ID, &friend.Name)
		if err != nil {
			return
		}
		friends = append(friends, friend)
	}
	return
}
