package lib

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func conn() *sqlx.DB {
	godotenv.Load()
	db, err := sqlx.Connect("postgres", "user=postgres dbname=go-coffee-shop password=1 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

var DB *sqlx.DB = conn()
