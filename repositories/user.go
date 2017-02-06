package repositories

import "framework/models"

type UserRepository interface {
	Create(*models.User) error
}

type UserRepositoryImpl struct {
}

func (repo *UserRepositoryImpl) Create(user *models.User) error {
	return db.Create(user).Error
}
