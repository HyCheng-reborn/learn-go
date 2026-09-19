package user

import "fmt"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUser(id int64) (*User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("get user %d: %w", id, err)
	}
	return user, nil
}
