package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionStringDb = ""
	ApiPort            = 0
	SecretKey []byte
)

func LoadEnv() {
	var erro error

	erro = godotenv.Load()
	if erro != nil {
		log.Fatal(erro)
	}

	ApiPort, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		ApiPort = 8000
	}

	ConnectionStringDb = fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=%s&loc=%s",
		os.Getenv("USER_DB"),
		os.Getenv("PASSWORD_DB"),
		os.Getenv("NAME_DB"),
		os.Getenv("CHARSET"),
		os.Getenv("PARSETIME"),
		os.Getenv("LOC"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
