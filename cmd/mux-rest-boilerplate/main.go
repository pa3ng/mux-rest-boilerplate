package main

import (
	"net/http"

	"github.com/pa3ng/mux-rest-boilerplate/internal/app"
)

func main() {
	router := app.NewRouter()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
