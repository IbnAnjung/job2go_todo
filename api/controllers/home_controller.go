package controllers

import (
	"net/http"

	"github.com/ibnanjung/job2go_todo/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, request *http.Request) {
	responses.JSON(w, http.StatusOK, "welcome")
}
