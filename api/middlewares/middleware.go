package middlewares

import (
	"errors"
	"net/http"

	"github.com/ibnanjung/job2go_todo/api/auth"
	"github.com/ibnanjung/job2go_todo/api/responses"
)

//SetMIddlewareJSON method
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, request)
	}
}

//SetMiddlewareAuthentication
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		err := auth.TokenValidation(request)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		next(w, request)
	}
}
