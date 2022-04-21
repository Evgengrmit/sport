package main

import (
	"log"
	"sport/pkg/handler"
	"sport/pkg/repository"
	"sport/pkg/service"
)

func RunServer() error {
	db, err := repository.GetConnection()
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
