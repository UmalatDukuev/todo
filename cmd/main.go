package main

import (
	"log"

	"github.com/UmalatDukuev/todo"
	"github.com/UmalatDukuev/todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	// if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
