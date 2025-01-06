package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var user models.User

	erro = json.Unmarshal(body, &user)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusBadRequest, erro)
		return
	}

	erro = user.Prepare("cadastro")
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	user.Id, erro = repository.Create(user)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JsonResponse(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNickOfUser := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := database.Connect()
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)

	users, erro := repository.Find(nameOrNickOfUser)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JsonResponse(w, http.StatusOK, users)
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 8)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	user, erro := repository.FindById(userId)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JsonResponse(w, http.StatusOK, user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 8)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusBadRequest, erro)
		return
	}

	body, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	erro = json.Unmarshal(body, &user)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusBadRequest, erro)
		return
	}

	erro = user.Prepare("edicao")
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	erro = repository.Update(userId, user)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JsonResponse(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 8)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	erro = repository.Delete(userId)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JsonResponse(w, http.StatusNoContent, nil)
}
