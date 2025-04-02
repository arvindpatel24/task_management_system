package task

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleCreateTask(w http.ResponseWriter, r *http.Request, useCase UseCase) {
	fmt.Println("Create task function called.")

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createdTask, err := useCase.CreateTask(task)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdTask)
}

func HandleListTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List tasks function called.")
	w.Write([]byte("All task list fetched."))
}

func HandleGetTaskById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get task function called.")
	w.Write([]byte("Task Fetched."))
}

func HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update task function called.")
	w.Write([]byte("Task Updated"))
}
func HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete task function called.")
	w.Write([]byte("Task Deleted"))
}
