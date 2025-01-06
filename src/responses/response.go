package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonError struct {
	Erro string `json:"erro"`
}

func JsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		erro := json.NewEncoder(w).Encode(data)
		if erro != nil {
			log.Fatal(erro)
		}
	}

}

func JsonErrorResponse(w http.ResponseWriter, statusCode int, erro error) {
	e := jsonError{Erro: erro.Error()}
	JsonResponse(w, statusCode, e)
}
