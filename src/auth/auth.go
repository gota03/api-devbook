package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint8) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Minute * 10).Unix()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.SecretKey)
}

func ValidateToken(r *http.Request) error {
	tokenString := extractTokenFromHeader(r)
	token, erro := jwt.Parse(tokenString, returnKeyOfAuth)
	if erro != nil {
		return erro
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func extractTokenFromHeader(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnKeyOfAuth(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("método de assinatura inesperado, %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := extractTokenFromHeader(r)
	token, erro := jwt.Parse(tokenString, returnKeyOfAuth)
	if erro != nil {
		return 0, erro
	}

	permissions, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userIdFormated := fmt.Sprintf("%0.f", permissions["userId"])
		userId, erro := strconv.ParseUint(userIdFormated, 10, 64)
		if erro != nil {
			return 0, erro
		}
		return userId, nil
	}
	return 0, errors.New("token inválido")
}
