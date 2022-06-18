package infrastructure

import (
	"database/sql"

	"problem1/domain/model"
	"problem1/domain/repository"
)

type UserRepositoryS struct {
	DB *sql.DB
}

// NewUserRepository はUserRepositorySのコンストラクタです。
func NewUserRepository(db *sql.DB) repository.UserRepositoryI {
	return &UserRepositoryS{DB: db}
}

// FindUser はUserをidで取得します。
func (ur *UserRepositoryS) FindUser(id int) (*model.User, error) {
	row, err := ur.DB.Query(`SELECT * FROM users WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var u model.User
	err = row.Scan(&u.ID, &u.UserID, &u.Name)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// FindUserList はUser listをidで取得します。
func (ur *UserRepositoryS) FindUserList() ([]*model.User, error) {
	rows, err := ur.DB.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	uList := make([]*model.User, 0)
	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.UserID, &u.Name)
		uList = append(uList, &u)
	}

	return uList, nil
}

// FindUserID はuser_idをidで取得します。
func (ur *UserRepositoryS) FindUserID(id int) (int, error) {
	row, err := ur.DB.Query(`SELECT user_id FROM users WHERE id = ?`, id)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var userID int
	row.Next()
	err = row.Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// FindID はidをuser_idで取得します。
func (ur *UserRepositoryS) FindID(uID int) (int, error) {
	row, err := ur.DB.Query(`SELECT id FROM users WHERE user_id = ?`, uID)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var id int
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// FindFriendLinkList はfriend link listをidで取得します。
func (ur *UserRepositoryS) FindFriendLinkList(id int) ([]*model.Link, error) {
	uID, err := ur.FindUserID(id)
	if err != nil {
		return nil, err
	}

	rows, err := ur.DB.Query(`SELECT * FROM friend_link WHERE user1_id = ? || user2_id = ?`, uID, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fl := make([]*model.Link, 0)
	for rows.Next() {
		var f model.Link
		rows.Scan(&f.ID, &f.User1ID, &f.User2ID)
		fl = append(fl, &f)
	}

	return fl, nil
}

// FindBlockList はblock listをidで取得します。
func (ur *UserRepositoryS) FindBlockList(id int) ([]*model.Link, error) {
	uID, err := ur.FindUserID(id)
	if err != nil {
		return nil, err
	}

	rows, err := ur.DB.Query(`SELECT * FROM block_list WHERE user1_id = ? || user2_id = ?`, uID, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bList := make([]*model.Link, 0)
	for rows.Next() {
		var b model.Link
		rows.Scan(&b.ID, &b.User1ID, &b.User2ID)
		bList = append(bList, &b)
	}

	return bList, nil
}
