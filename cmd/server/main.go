package main

import (
	"log"
	"os"
	"sport/pkg/handler"
	"sport/pkg/repository"
	"sport/pkg/service"
)

func RunServer() error {
	host := os.Getenv("HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")
	db, err := repository.ConnectPostgresDB(repository.Config{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbname,
		SSLMode:  sslmode,
	})
	if err != nil {
		log.Fatalf("error occurred when initialization database: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	serv := service.NewService(repos)
	h := handler.NewHandler(serv)
	r := h.InitRoutes()
	return r.Run()
}

func main() {

	err := RunServer()
	if err != nil {
		log.Fatalln(err.Error())
	}
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
