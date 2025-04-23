package api

import (
	"github.com/gorilla/mux"
	"github.com/pvnptl/task-service/internal/api/handlers"
)

// NewRouter sets up the routes for task APIs
func NewRouter(taskHandler *handlers.TaskHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")     // Add a task
	router.HandleFunc("/tasks", taskHandler.ListTasks).Methods("GET")       // Get all tasks (with filter/pagination)
	router.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")    // Get one task by ID
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT") // Update a task
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE") // Delete a task

	return router
}
