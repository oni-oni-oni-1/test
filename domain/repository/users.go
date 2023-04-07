package repository

import (
	"gaudy_code/domain/model"
)

type UserRepository interface {
	FindUser(userID int64) (*model.User, error)
	UpdateUser(user model.User) error
}
