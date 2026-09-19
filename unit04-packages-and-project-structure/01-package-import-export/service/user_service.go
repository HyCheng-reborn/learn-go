package service

import (
	"fmt"

	"learn-go/unit04-packages-and-project-structure/01-package-import-export/model"
	"learn-go/unit04-packages-and-project-structure/01-package-import-export/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int64) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("get user %d: %w", id, err)
	}
	return user, nil
}
