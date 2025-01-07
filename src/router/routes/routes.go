package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri                   string
	Method                string
	Func                  func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.RequireAuthentication {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Func))).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}
	return r
}
