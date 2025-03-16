package server

import (
	"github.com/guneyin/todo-app/todo-api/handler"
	"log"
	"net/http"
)

func StartServer(port string) error {
	log.Println("Starting server...")

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	return http.ListenAndServe(":"+port, mux)
}
