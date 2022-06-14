package repository

import (
	"problem1/domain/model"
)

type UserRepository interface {
	// Create(user *model.User) error
	FindByID(id int) (*model.User, error)
	FindUserIDByID(id int) (int, error)
	FindIDByUserID(userID int) (int, error)
	FindFriendsByID(id int) ([]*model.Link, error)
	// Delete(user *model.User) error
}
