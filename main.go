package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	apiPort := fmt.Sprintf(":%d", config.ApiPort)
	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(apiPort, r))
}
