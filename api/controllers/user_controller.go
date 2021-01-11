package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ibnanjung/job2go_todo/api/models"
	"github.com/ibnanjung/job2go_todo/api/responses"
)

//CreateUser method
func (server *Server) CreateUser(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("cant creating user"))
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", request.Host, request.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)

}
