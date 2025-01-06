package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, erro := database.Connect()
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	userSaveInDb, erro := repository.FindByEmail(user.Email)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	erro = security.ComparePassword(userSaveInDb.Password, user.Password)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusUnauthorized, erro)
		return
	}

	token, _ := auth.CreateToken(user.Id)
	w.Write([]byte(token))
}
