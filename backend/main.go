package main

import (
	"Task-scheduler/handlers"
	_ "Task-scheduler/handlers"
	"fmt"
	"net/http"
)

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//})
	http.HandleFunc("/tasks", handlers.GetTask)
	http.HandleFunc("/tasks/create", handlers.CreateTask)

	fmt.Println("Listening on port 5173", "http://localhost:5173")
	http.ListenAndServe(":5173", nil)
}
