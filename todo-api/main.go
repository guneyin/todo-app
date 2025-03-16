package main

import (
	"log"
	"os"

	"github.com/guneyin/todo-app/todo-api/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Fatal(server.StartServer(os.Getenv("API_PORT")))
}
