package repository

import (
	"problem1/domain/model"
)

type UserRepository interface {
	FindUser(id int) (*model.User, error)
	FindUserID(id int) (int, error)
	FindID(userID int) (int, error)
	FindFriendLinkList(id int) ([]*model.Link, error)
	FindBlockList(id int) ([]*model.Link, error)
}
