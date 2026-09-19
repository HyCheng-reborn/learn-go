package repository

import (
	"learn-go/unit04-packages-and-project-structure/01-package-import-export/model"
)

type MemoryUserRepository struct {
	users map[int64]*model.User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: map[int64]*model.User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
	}
}

func (r *MemoryUserRepository) FindByID(id int64) (*model.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}
