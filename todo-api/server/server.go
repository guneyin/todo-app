package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/guneyin/todo-app/todo-api/handler"
)

func StartServer(port string) error {
	log.Println("Starting server on port", port)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		Handler:           mux,
	}

	return server.ListenAndServe()
}
