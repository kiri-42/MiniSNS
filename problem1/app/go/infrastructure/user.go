package infrastructure

import (
	"database/sql"

	"problem1/domain/model"
	"problem1/domain/repository"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) FindByID(id int) (*model.User, error) {
	rows, err := ur.DB.Query(`SELECT * FROM users WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	rows.Next()
	user := new(model.User)
	err = rows.Scan(&(user.ID), &(user.UserID), &(user.Name))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) FindUserIDByID(id int) (int, error) {
	row, err := ur.DB.Query(`SELECT user_id FROM users WHERE id = ?`, id)
	if err != nil {
		return 0, err
	}

	var userID int
	row.Next()
	err = row.Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (ur *UserRepository) FindFriendsByID(id int) ([]*model.Link, error) {
	userID, err := ur.FindUserIDByID(id)
	if err != nil {
		return nil, err
	}

	rows, err := ur.DB.Query(`SELECT * FROM friend_link WHERE user1_id = ? || user2_id = ?`, userID, userID)
	if err != nil {
		return nil, err
	}

	friends := make([]*model.Link, 0)
	for rows.Next() {
		var friend model.Link
		rows.Scan(&friend.ID, &friend.User1ID, &friend.User2ID)
		friends = append(friends, &friend)
	}

	return friends, nil
}
