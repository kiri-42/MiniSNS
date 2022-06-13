package repository

import (
	"problem1/domain/model"
)

type UserRepository interface {
	// Create(user *model.User) error
	FindByID(id int) (*model.User, error)
	// Delete(user *model.User) error
}
