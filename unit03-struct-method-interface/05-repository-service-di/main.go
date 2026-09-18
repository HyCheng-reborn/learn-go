package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int64
	Name string
}

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	FindByID(id int64) (*User, error)
}

type MemoryUserRepository struct {
	users map[int64]*User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: map[int64]*User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
	}
}

func (r *MemoryUserRepository) FindByID(id int64) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int64) (*User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("get user %d: %w", id, err)
	}
	return user, nil
}

func main() {
	repo := NewMemoryUserRepository()
	service := NewUserService(repo)

	user, err := service.GetUser(1)
	if err != nil {
		fmt.Println("get user failed:", err)
		return
	}
	fmt.Println("found:", user.Name)

	_, err = service.GetUser(999)
	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("user 999 was not found")
	}
}
