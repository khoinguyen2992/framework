package users

import "framework/models"

type UserRepository interface {
	Create(*models.User) error
}
