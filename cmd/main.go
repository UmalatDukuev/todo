package main

import (
	"log"
	"todo"
	"todo/pkg/handler"
)

func main() {
	handler := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
