package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/ibnanjung/job2go_todo/api/auth"
	"github.com/ibnanjung/job2go_todo/api/models"
	"github.com/ibnanjung/job2go_todo/api/responses"
	"golang.org/x/crypto/bcrypt"
)

//Login method
func (server *Server) Login(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("password and usernamenot matched"))
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

//SignIn method
func (server *Server) SignIn(username, password string) (string, error) {
	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("username= ?", username).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(user.ID)
}
