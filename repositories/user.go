package repositories

import "framework/models"

type UserRepository struct {
}

func (repo *UserRepository) Create(user *models.User) error {
	return db.Create(user).Error
}
