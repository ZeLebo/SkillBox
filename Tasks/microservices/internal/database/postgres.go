package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"user/internal/domain"
	u "user/internal/user"
)

type Client struct {
	db     *gorm.DB
	logger *log.Logger
	*domain.DBConfig
}

func NewClient(cfg *domain.DBConfig, logger *log.Logger) (*Client, error) {
	client := &Client{
		DBConfig: cfg,
		logger:   logger,
	}
	db, err := client.Connect()
	if err != nil {
		return nil, err
	}
	client.db = db
	return client, nil
}

func (c *Client) dsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		c.Host, c.Port, c.Username, c.Password, c.DBName)
}

func (c *Client) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.dsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&Model{}); err != nil {
		return nil, fmt.Errorf("cannot migrate: %w", err)
	}
	if err = db.AutoMigrate(&Friend{}); err != nil {
		return nil, fmt.Errorf("cannot migrate: %w", err)
	}
	return db, nil
}

func (c *Client) GetUsers() ([]u.User, error) {
	var users []Model
	if err := c.db.Find(&users).Error; err != nil {
		return nil, err
	}
	var result []u.User
	for _, user := range users {
		result = append(result, u.User{
			ID:   user.ID,
			Name: user.Name,
			Age:  user.Age,
		})
	}
	return result, nil
}

func (c *Client) CreateUser(model Model) (int, error) {
	if err := c.db.Create(&model).Error; err != nil {
		return 0, err
	}
	var id int
	if err := c.db.Model(&Model{}).Where("name = ? AND age = ?", model.Name, model.Age).Select("id").First(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (c *Client) DeleteUser(id int) error {
	// delete from models the record with id = id
	if err := c.db.Delete(&Model{}, id).Error; err != nil {
		return err
	}
	// delete all the relations in friends table
	if err := c.db.Where("friend1 = ? OR friend2 = ?", id, id).Delete(&Friend{}).Error; err != nil {
		return err
	}
	return nil
}

func (c *Client) ChangeAge(id, age int) error {
	if err := c.db.Model(&Model{}).Where("id = ?", id).Update("age", age).Error; err != nil {
		return err
	}
	return nil
}

func (c *Client) MakeFriends(id1, id2 int) error {
	// add id1 to id2 friends
	if err := c.db.Create(&Friend{
		First:  id1,
		Second: id2,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (c *Client) GetFriends(id int) ([]u.User, error) {
	var friendsIdFirst []int

	if err := c.db.Model(&Friend{}).Where("friend1 = ?", id).Select("friend2").Find(&friendsIdFirst).Error; err != nil {
		return nil, err
	}
	var friendsIdSecond []int
	// check all the values where second is id
	if err := c.db.Model(&Friend{}).Where("friend2 = ?", id).Select("friend1").Find(&friendsIdSecond).Error; err != nil {
		return nil, err
	}
	type null struct{}
	friendsId := make(map[int]null)
	for _, id := range friendsIdFirst {
		friendsId[id] = null{}
	}
	for _, id := range friendsIdSecond {
		friendsId[id] = null{}
	}

	var friends []u.User
	for id := range friendsId {
		newFriend, err := c.GetUserByID(id)
		if err == nil {
			friends = append(friends, newFriend)
		}
	}
	return friends, nil
}

func (c *Client) CheckUser(user Model) bool {
	var model Model
	if err := c.db.Where("name = ? AND age = ?", user.Name, user.Age).First(&model).Error; err != nil {
		return false
	}
	return true
}

func (c *Client) GetUserByID(id int) (u.User, error) {
	var user Model
	if err := c.db.First(&user, id).Error; err != nil {
		return u.User{}, err
	}
	return u.User{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}, nil
}

func (c *Client) GetUserID(user Model) int {
	var id int
	if err := c.db.Model(&Model{}).Where("name = ? AND age = ?", user.Name, user.Age).Select("id").First(&id).Error; err != nil {
		return 0
	}
	return id
}
