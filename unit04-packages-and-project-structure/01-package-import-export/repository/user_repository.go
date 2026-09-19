package repository

import (
	"errors"
	"learn-go/unit04-packages-and-project-structure/01-package-import-export/model"
)

type UserRepository interface {
	FindByID(id int64) (*model.User, error)
}

var ErrUserNotFound = errors.New("user not found")
