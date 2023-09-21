package repository

import (
	"AnaUserService/model"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := u.db.QueryRow("SELECT id, email, phone, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Phone)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetUserByEmailOrPhone(username string) (*model.User, error) {
	var user model.User
	err := u.db.QueryRow("SELECT id, email, phone, password FROM users WHERE email = $1 OR phone = $1", username).Scan(&user.ID, &user.Email, &user.Phone, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
