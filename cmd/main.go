package main

import (
	"os" // work with OS

	"github.com/joho/godotenv"                        // env config
	_ "github.com/lib/pq"                             // postgres driver
	"github.com/sirupsen/logrus"                      // log formatter (JSON chosen below)
	"github.com/spf13/viper"                          // viper config
	"github.com/yarikTri/todo-app"                    // root github repo
	"github.com/yarikTri/todo-app/package/handler"    // package handler
	"github.com/yarikTri/todo-app/package/repository" // package repository
	"github.com/yarikTri/todo-app/package/service"    // package service
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error while loading env vars: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize DB: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
