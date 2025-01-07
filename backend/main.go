package main

import (
	"Task-scheduler/handlers"
	_ "Task-scheduler/handlers"
	_ "database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//})
	http.HandleFunc("/tasks", handlers.GetTask)
	http.HandleFunc("/tasks/create", handlers.CreateTask)

	fmt.Println("Listening on port 8080", "http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	//env file loading
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conn = ConnectDB()
	defer CloseDB()
}
