package main

import (
	_ "github.com/lib/pq"
	"log"
	app "rest_api"
	"rest_api/internal/handler"
	"rest_api/internal/repository"
	"rest_api/internal/service"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	server := app.NewServer()
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Сервер сломался %s", err.Error())
	}
}
