package tabledriventests

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")
var ErrDataBase = errors.New("database error")

type User struct {
	ID   int64
	Name string
}

type Repository interface {
	FindByID(id int64) (*User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUser(id int64) (*User, error) {
	usr, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("get user %d: %w", id, err)
	}
	return usr, nil
}
