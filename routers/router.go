package routers

import (
	"framework/handlers"
	"framework/handlers/users"
	"framework/repositories"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	logger := logrus.New()

	r := mux.NewRouter()

	r.Handle("/", handlers.HelloHandler{}).Methods("GET")

	r.Handle("/users/", users.CreateUserHandler{&repositories.UserRepositoryImpl{}, logger}).Methods("POST")

	return &Router{r}
}
