package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// JSON file structure : ID, TITLE , DESCRIPTION
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var tasks = []Task{}

// getting tasks: modify to get from db
func GetTask(w http.ResponseWriter, r *http.Request) {
	rows, err := conn.Query(context.Background(), "SELECT id , title, description FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		error := rows.Scan(&task.ID, &task.Title, &task.Description)
		if error != nil {
			log.Println(error) // Log the error and continue
			continue
		}
		tasks = append(tasks, task)
	}

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
