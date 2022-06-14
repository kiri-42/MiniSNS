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





