package main

import (
	"errors"
	"fmt"
)

// 闭卷练习：从零实现 unit03/05 的 repository → service → DI 链路。
// 要求见本目录 notes.md。写之前不要回看 05-repository-service-di/main.go。
// 全部代码打在这个文件里（单 package main，不拆包——拆包是 unit04 的事）。

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
		users: map[int64]*User{1: {ID: 1, Name: "Alice"}},
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

type MockUserRepository struct {
	knownIDs map[int64]bool
}

func (r *MockUserRepository) FindByID(id int64) (*User, error) {
	if !r.knownIDs[id] {
		return nil, ErrUserNotFound
	}
	return &User{ID: id, Name: "MockUser"}, nil
}

func main() {
	service := NewUserService(NewMemoryUserRepository())
	for _, id := range []int64{1, 999} {
		user, err := service.GetUser(id)
		switch {
		case errors.Is(err, ErrUserNotFound):
			fmt.Printf("user %d was not found\n", id)
			fmt.Printf("wrapped: %v\n", err)
		case err != nil:
			fmt.Printf("error: %v\n", err)
		default:
			fmt.Printf("found: %s\n", user.Name)
		}
	}

	mockService := NewUserService(&MockUserRepository{
		knownIDs: map[int64]bool{100: true},
	})
	for _, id := range []int64{100, 777} {
		user, err := mockService.GetUser(id)
		switch {
		case errors.Is(err, ErrUserNotFound):
			fmt.Printf("mock: user %d was not found\n", id)
		case err != nil:
			fmt.Printf("mock error: %v\n", err)
		default:
			fmt.Printf("found: %s\n", user.Name)
		}
	}
}
