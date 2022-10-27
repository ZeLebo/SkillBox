package service

import (
	"log"
	"user/internal/database"
	"user/internal/domain"
	u "user/internal/user"
)

type Service struct {
	client ClientInterface
	logger *log.Logger
}

func NewService(client *database.Client, logger *log.Logger) *Service {
	return &Service{
		client: client,
		logger: logger,
	}
}

// GetAllUsers func to return all the users in the map
func (s *Service) GetAllUsers(_ *domain.Request) ([]u.User, error) {
	users, err := s.client.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Create function returns id of user
func (s *Service) Create(req *domain.Request) (int, error) {
	user := u.User{
		Name: req.Name,
		Age:  req.Age,
	}

	id, err := s.client.CreateUser(user)
	if err != nil {
		return 0, err
	}
	// put the friends in the database if they are not exist
	for _, friend := range req.Friends {
		user = u.User{
			Name: friend.Name,
			Age:  friend.Age,
		}

		friendId := s.client.GetUserID(user)
		if friendId == 0 {
			friendId, err = s.client.CreateUser(user)
			if err != nil {
				return 0, err
			}
		}
		// add id and friend id to the friend table
		err = s.client.MakeFriends(id, friendId)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

// ChangeAge to change the age of specific user
func (s *Service) ChangeAge(req *domain.Request) error {
	err := s.client.ChangeAge(req.TargetID, req.Age)
	if err != nil {
		return err
	}
	return nil
}

// GetFriends of specific user
func (s *Service) GetFriends(req *domain.Request) ([]u.User, error) {
	// return the list of friends
	friends, err := s.client.GetFriends(req.TargetID)
	if err != nil {
		return nil, err
	}
	return friends, nil
}

// MakeFriends make friends from 2 users
func (s *Service) MakeFriends(req *domain.Request) error {
	// add id1 to id2 friends
	err := s.client.MakeFriends(req.TargetID, req.SourceID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUser(req *domain.Request) error {
	err := s.client.DeleteUser(req.TargetID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CheckUser(req *domain.Request) bool {
	user := u.User{
		Name: req.Name,
		Age:  req.Age,
	}
	return s.client.CheckUser(user)
}

func (s *Service) GetUserByID(id int) (u.User, error) {
	user, err := s.client.GetUserByID(id)
	if err != nil {
		return u.User{}, err
	}
	return user, nil
}
