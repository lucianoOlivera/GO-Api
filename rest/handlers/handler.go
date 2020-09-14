package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

//GetUsers obtiene todos los usuarios
func GetUsers(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	if user, err := GetUserId(r); err != nil {
		models.SendNotFoud(w)

	} else {
		models.SendData(w, user)
	}

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendUnProcessableEntity(w)
	} else {
		models.SendData(w, models.SaveUser(user))
	}
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if user, err := GetUserId(r); err != nil {
		models.SendNotFoud(w)
	} else {
		models.DeleteUser(user.Id)
		models.SendNoContent(w)
	}
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserId(r)
	if err != nil {
		models.SendNotFoud(w)
	}
	return
	userResponse := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userResponse); err != nil {
		models.SendUnProcessableEntity(w)

		return
	}
	user = models.UpdateUser(user, userResponse.Name, userResponse.Password)
	models.SendData(w, user)
}

func GetUserId(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	UserId, _ := strconv.Atoi(vars["id"])
	if user, err := models.GetUserId(UserId); err != nil {
		return user, err

	} else {
		return user, nil
	}

}
