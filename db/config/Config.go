package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDB() *sql.DB {

	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatal("Error loading env file")
	// }
	dbConfig, err := sql.Open("postgres", "user=anupam password=mailpass dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal("error loading database config", err)
	}
	return dbConfig
}
