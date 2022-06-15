package repository

import (
	"problem1/domain/model"
)

type UserRepositoryI interface {
	FindID(userID int) (int, error)
	FindUserID(id int) (int, error)
	FindUser(id int) (*model.User, error)
	FindUserList() ([]*model.User, error)
	FindFriendLinkList(id int) ([]*model.Link, error)
	FindBlockList(id int) ([]*model.Link, error)
}
