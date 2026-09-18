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

// MockUserRepository 只认识 knownIDs 里的 id，其余一律返回 ErrUserNotFound，
// 这样不接真实数据源也能同时演示成功与失败两条路径。
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
	repo := NewMemoryUserRepository()
	service := NewUserService(repo)

	user, err := service.GetUser(1)
	if err != nil {
		fmt.Println("get user failed:", err)
		return
	}
	fmt.Println("found:", user.Name)

	// errors.Is(nil, ErrUserNotFound) 返回 false，所以还要留一个 err != nil 分支，
	// 否则「不是 NotFound 的其他错误」会被静默吞掉。
	_, err = service.GetUser(999)
	switch {
	case errors.Is(err, ErrUserNotFound):
		fmt.Println("user 999 was not found")
		fmt.Println("wrapped:", err)
	case err != nil:
		fmt.Println("unexpected error:", err)
	}

	// repo 的静态类型是 *MemoryUserRepository（:= 推断出的是具体类型，不是接口），
	// 所以不能再把 *MockUserRepository 赋给它。改用独立变量：具体类型在传给
	// NewUserService(repo UserRepository) 的形参时完成隐式接口转换。
	mockRepo := &MockUserRepository{knownIDs: map[int64]bool{100: true}}
	mockService := NewUserService(mockRepo)

	user, err = mockService.GetUser(100)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("found:", user.Name)

	// mock 也能驱动失败路径，这正是把依赖抽成接口的价值所在。
	_, err = mockService.GetUser(777)
	switch {
	case errors.Is(err, ErrUserNotFound):
		fmt.Println("mock: user 777 was not found")
		fmt.Println("mock wrapped:", err)
	case err != nil:
		fmt.Println("mock unexpected error:", err)
	}
}
