package main

import (
	"log"

	todo "github.com/AndriyAntonenko/todoApp"
	"github.com/AndriyAntonenko/todoApp/pkg/handler"
	"github.com/AndriyAntonenko/todoApp/pkg/repository"
	"github.com/AndriyAntonenko/todoApp/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
