package users

import (
	"framework/forms"
	"framework/handlers"
	"framework/models"
	"net/http"

	"github.com/mholt/binding"
)

type CreateUserHandler struct {
	UserRepo UserRepository
}

func (handler CreateUserHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	createUserForm := new(forms.CreateUserForm)
	errs := binding.Bind(req, createUserForm)
	if errs.Len() > 0 {
		es := handlers.Errors{errs}
		handlers.RenderError(rw, http.StatusBadRequest, es.GetErrors())
		return
	}

	user := models.User(*createUserForm)
	if err := handler.UserRepo.Create(&user); err != nil {
		handlers.RenderError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	handlers.RenderJSON(rw, http.StatusCreated, user)
}
