package forms

import (
	"net/http"
	"net/mail"

	"github.com/mholt/binding"
)

type CreateUserForm struct {
	Id    int64  `json:"-"`
	Email string `json:"email"`
}

func (form *CreateUserForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&form.Email: binding.Field{
			Form:     "email",
			Required: true,
		},
	}
}

func (form *CreateUserForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if _, err := mail.ParseAddress(form.Email); err != nil {
		errs = append(errs, binding.Error{
			FieldNames:     []string{"email"},
			Classification: binding.RequiredError,
			Message:        "Invalid",
		})
	}

	return errs
}
