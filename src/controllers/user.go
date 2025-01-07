package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
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
		if erro.Error() == "usuário não encontrado" {
			responses.JsonErrorResponse(w, http.StatusNotFound, erro)
			return
		}
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JsonResponse(w, http.StatusOK, user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusBadRequest, erro)
		return
	}

	userIdToken, erro := auth.ExtractUserId(r)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusUnauthorized, erro)
		return
	}

	if userId != userIdToken {
		errForbidden := errors.New("não é possível atualizar um usuário que não seja o seu")
		responses.JsonErrorResponse(w, http.StatusForbidden, errForbidden)
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

	userIdToken, erro := auth.ExtractUserId(r)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusUnauthorized, erro)
		return
	}

	if userId != userIdToken {
		errForbidden := errors.New("não é possível deletar um usuário que não seja o seu")
		responses.JsonErrorResponse(w, http.StatusForbidden, errForbidden)
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

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, erro := auth.ExtractUserId(r)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusForbidden, erro)
		return
	}

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusBadRequest, erro)
		return
	}

	if followerId == userId {
		errForbidden := errors.New("não é possível seguir você mesmo")
		responses.JsonErrorResponse(w, http.StatusForbidden, errForbidden)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	erro = repository.Follow(userId, followerId)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JsonResponse(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, erro := auth.ExtractUserId(r)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusForbidden, erro)
		return
	}

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusBadRequest, erro)
		return
	}

	if followerId == userId {
		errForbidden := errors.New("não é possível para de seguir você mesmo")
		responses.JsonErrorResponse(w, http.StatusForbidden, errForbidden)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	erro = repository.Unfollow(userId, followerId)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JsonResponse(w, http.StatusNoContent, nil)
}

func SearchFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
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
	followers, erro := repository.SearchFollowers(userId)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JsonResponse(w, http.StatusOK, followers)
}

func SearchFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
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
	followed, erro := repository.SearchFollowing(userId)
	if erro != nil {
		responses.JsonErrorResponse(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JsonResponse(w, http.StatusOK, followed)
}
