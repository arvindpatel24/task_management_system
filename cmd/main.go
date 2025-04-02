package main

import (
	"fmt"
	"net/http"

	"github.com/arvindpatel24/task_management_system/internal/task"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// Set up the HTTP router
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	// routes
	router.Post("/tasks", task.HandleCreateTask)
	router.Get("/tasks", task.HandleListTasks)
	router.Get("/tasks/{id}", task.HandleGetTaskById)
	router.Put("/tasks/{id}", task.HandleUpdateTask)
	router.Delete("/tasks{id}", task.HandleDeleteTask)

	fmt.Println("Starting server...")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}
}
