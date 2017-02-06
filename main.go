package main

import (
	"framework/repositories"
	"framework/routers"
	"net/http"
)

func main() {
	if err := repositories.InitDatabase(); err != nil {
		panic(err)
	}
	defer repositories.CloseDatabase()

	router := routers.NewRouter()

	http.ListenAndServe(":8080", router)
}
