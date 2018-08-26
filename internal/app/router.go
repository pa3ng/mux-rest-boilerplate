package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const APIVersionPath = "/api/v1"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		route("/"),
		Index,
	},
	Route{
		"Health",
		"GET",
		route("/health"),
		Health,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		log.Printf("%s: %s", route.Method, route.Pattern)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func route(p string) string {
	return APIVersionPath + p
}
