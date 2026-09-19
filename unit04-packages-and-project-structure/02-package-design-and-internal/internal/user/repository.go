package user

import "errors"

var ErrNotFound = errors.New("user not found")

type Repository interface {
	FindByID(id int64) (*User, error)
}

type MemoryRepository struct {
	users map[int64]*User
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users: map[int64]*User{
			1: {ID: 1, Name: "Alice"},
		},
	}
}

func (r *MemoryRepository) FindByID(id int64) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, ErrNotFound
	}
	return user, nil
}

func (r *MemoryRepository) Count() int {
	return len(r.users)
}
