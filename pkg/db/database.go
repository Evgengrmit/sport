package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

var DB *sqlx.DB

func ConnectPostgresDB(cfg Config) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}
	DB = db
	fmt.Println("Connection started")
}
