package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/yarikTri/todo-app"
	"github.com/yarikTri/todo-app/package/handler"
	"github.com/yarikTri/todo-app/package/repository"
	"github.com/yarikTri/todo-app/package/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing configs: %s", err.Error())
	}

	repository := repository.NewRepository()
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
