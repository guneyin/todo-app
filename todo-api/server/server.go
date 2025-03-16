package server

import (
	"log"
	"net/http"
)

func StartServer(port string) error {
	log.Println("Starting server...")

	mux := http.NewServeMux()

	return http.ListenAndServe(":"+port, mux)
}
