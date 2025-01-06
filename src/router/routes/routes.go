package routes

import (
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
		r.HandleFunc(route.Uri, route.Func).Methods(route.Method)
	}
	return r
}
