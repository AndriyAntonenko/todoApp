package main

import (
	"log"

	todo "github.com/AndriyAntonenko/todoApp"
	"github.com/AndriyAntonenko/todoApp/pkg/handler"
	"github.com/AndriyAntonenko/todoApp/pkg/repository"
	"github.com/AndriyAntonenko/todoApp/pkg/service"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
