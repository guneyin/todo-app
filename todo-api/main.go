package main

import (
	"github.com/guneyin/todo-app/todo-api/server"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	log.Fatal(server.StartServer(os.Getenv("PORT")))
}
