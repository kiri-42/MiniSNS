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

// // 未実装
// func (ur *UserRepository) Create(user *model.User) error {
// 	return nil
// }

func (ur *UserRepository) FindByID(id int) (*model.User, error) {
	rows, err := ur.DB.Query(`SELECT * FROM users WHERE user_id = ?`, id)
	if err != nil {
		return nil, err
	}

	rows.Next()
	user := new(model.User)
	err = rows.Scan(user.ID, user.UserID, user.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// func (ur *UserRepository) Delete(user *model.User) error {
// 	return nil
// }



