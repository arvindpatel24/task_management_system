package task

import (
	"fmt"
	"net/http"
)

func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create task function called.")
	w.Write([]byte("Task Created."))
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
