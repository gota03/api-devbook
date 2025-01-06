package database

import (
	"api/src/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConnectionStringDb)
	if erro != nil {
		log.Fatal(erro)
		return nil, erro
	}

	erro = db.Ping()
	if erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
