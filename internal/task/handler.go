package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func HandleListTasks(w http.ResponseWriter, r *http.Request, useCase UseCase) {
	fmt.Println("List tasks function called.")
	tasks, err := useCase.GetTasks()
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func HandleGetTaskById(w http.ResponseWriter, r *http.Request, useCase UseCase) {
	fmt.Println("Get task function called.")
	idStr := r.URL.Path[len("/tasks/"):]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := useCase.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func HandleUpdateTask(w http.ResponseWriter, r *http.Request, useCase UseCase) {
	fmt.Println("Update task function called.")
	idStr := r.URL.Path[len("/tasks/"):]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedTask, err := useCase.UpdateTask(id, task)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedTask)
}
func HandleDeleteTask(w http.ResponseWriter, r *http.Request, useCase UseCase) {
	fmt.Println("Delete task function called.")
	idStr := r.URL.Path[len("/tasks/"):]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = useCase.DeleteTask(id)
	if err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
