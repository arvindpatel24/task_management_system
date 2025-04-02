package main

import (
	"fmt"
	"net/http"

	"github.com/arvindpatel24/task_management_system/config"
	"github.com/arvindpatel24/task_management_system/internal/storage"
	"github.com/arvindpatel24/task_management_system/internal/task"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// Load Config
	config := config.LoadConfig()

	// Set up the DB Connection
	dbConnection := storage.NewDB(config.SqlAddress)

	// Initialize repository and use case layers
	repo := task.NewTaskRepository(dbConnection)
	useCase := task.NewTaskUseCase(repo)

	// Set up the HTTP router
	router := chi.NewRouter()

	// Basic middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	// routes
	router.Post("/tasks", func(w http.ResponseWriter, r *http.Request) { task.HandleCreateTask(w, r, useCase) })
	router.Get("/tasks", task.HandleListTasks)
	router.Get("/tasks/{id}", task.HandleGetTaskById)
	router.Put("/tasks/{id}", task.HandleUpdateTask)
	router.Delete("/tasks{id}", task.HandleDeleteTask)

	fmt.Println("Starting server...")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.ServerPort),
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}
}
