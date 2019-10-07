package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route blablabla...
type Route struct {
	nombre        string
	tipo          string
	path          string
	nombreFuncion http.HandlerFunc
}

// Routes blablabla
type Routes []Route

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.tipo).
			Name(route.nombre).
			Path(route.path).
			HandlerFunc(route.nombreFuncion)
	}
	return router
}

var routes = Routes{
	Route{
		"scrapData",
		"GET",
		"/scrapData",
		scrapData,
	},
}
