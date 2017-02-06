package repositories

import (
	"framework/models"
	"testing"
)

func TestCreateUser(t *testing.T) {
	userRepo := UserRepositoryImpl{}
	user := models.User{
		Email: "test@example.com",
	}

	if err := userRepo.Create(&user); err != nil {
		t.Errorf("It should NOT get error: %s", err.Error())
	}

	if user.Id == 0 {
		t.Errorf("It should persist user")
	}
}
