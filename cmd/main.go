package main

import (
	"log"

	"github.com/yarikTri/todo-app"
	"github.com/yarikTri/todo-app/package/handler"
	"github.com/yarikTri/todo-app/package/repository"
	"github.com/yarikTri/todo-app/package/service"
)

func main() {
	repository := repository.NewRepository()
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server %s", err.Error())
	}
}
