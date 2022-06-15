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

func (ur *UserRepository) FindUser(id int) (*model.User, error) {
	row, err := ur.DB.Query(`SELECT * FROM users WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	row.Next()
	u := new(model.User)
	err = row.Scan(&(u.ID), &(u.UserID), &(u.Name))
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) FindUserID(id int) (int, error) {
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

func (ur *UserRepository) FindID(userID int) (int, error) {
	row, err := ur.DB.Query(`SELECT id FROM users WHERE user_id = ?`, userID)
	if err != nil {
		return 0, err
	}

	var id int
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ur *UserRepository) FindFriendLinkList(id int) ([]*model.Link, error) {
	uID, err := ur.FindUserID(id)
	if err != nil {
		return nil, err
	}

	rows, err := ur.DB.Query(`SELECT * FROM friend_link WHERE user1_id = ? || user2_id = ?`, uID, uID)
	if err != nil {
		return nil, err
	}

	fl := make([]*model.Link, 0)
	for rows.Next() {
		var f model.Link
		rows.Scan(&f.ID, &f.User1ID, &f.User2ID)
		fl = append(fl, &f)
	}

	return fl, nil
}

func (ur *UserRepository) FindBlockList(id int) ([]*model.Link, error) {
	uID, err := ur.FindUserID(id)
	if err != nil {
		return nil, err
	}

	rows, err := ur.DB.Query(`SELECT * FROM block_list WHERE user1_id = ? || user2_id = ?`, uID, uID)
	if err != nil {
		return nil, err
	}

	bList := make([]*model.Link, 0)
	for rows.Next() {
		var b model.Link
		rows.Scan(&b.ID, &b.User1ID, &b.User2ID)
		bList = append(bList, &b)
	}

	return bList, nil
}
