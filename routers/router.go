package routers

import (
	"framework/handlers"
	"framework/handlers/users"
	"framework/repositories"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	r := mux.NewRouter()

	r.Handle("/", handlers.HelloHandler{}).Methods("GET")

	r.Handle("/users/", users.CreateUserHandler{&repositories.UserRepository{}}).Methods("POST")

	return &Router{r}
}
