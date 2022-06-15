package repository

import (
	"problem1/domain/model"
)

type UserRepository interface {
	FindByID(id int) (*model.User, error)
	FindUserIDByID(id int) (int, error)
	FindIDByUserID(userID int) (int, error)
	FindFriendsByID(id int) ([]*model.Link, error)
	FindBlockList(id int) ([]*model.Link, error)
}
