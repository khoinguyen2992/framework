package main

import (
	"framework/repositories"
	"framework/routers"
	"net/http"

	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
)

func main() {
	if err := repositories.InitDatabase(); err != nil {
		panic(err)
	}
	defer repositories.CloseDatabase()

	router := routers.NewRouter()

	n := negroni.New()
	n.Use(negronilogrus.NewMiddleware())
	n.UseHandler(router)

	http.ListenAndServe(":8080", n)
}
