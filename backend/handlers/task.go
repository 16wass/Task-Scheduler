package handlers

import (
	"encoding/json"
	"net/http"
)

// JSON file structure : ID , TITLE , DESCRIPTION
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var tasks = []Task{}

// getting tasks
func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// creating tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	tasks = append(
		tasks,
		task,
	)
	w.WriteHeader(http.StatusCreated)
}
