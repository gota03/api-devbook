package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"log"
	"net/http"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.RequestURI, r.Method, r.Host)
		nextFunc(w, r)
	}
}

func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		erro := auth.ValidateToken(r)
		if erro != nil {
			responses.JsonErrorResponse(w, http.StatusUnauthorized, erro)
			return
		}
		nextFunc(w, r)
	}
}
