package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pvnptl/task-service/internal/models"
	"github.com/pvnptl/task-service/internal/service"
)

// TaskHandler handles HTTP requests related to tasks
type TaskHandler struct {
	service service.TaskService
}

// NewTaskHandler creates and returns a new TaskHandler
func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// CreateTask handles POST /tasks — creates a new task
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	// Decode JSON body into task struct
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Save task using service
	err = h.service.CreateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with created task
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// GetTask handles GET /tasks/{id} — returns a task by ID
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// UpdateTask handles PUT /tasks/{id} — updates a task
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task.ID = id // Set task ID from URL path

	err = h.service.UpdateTask(&task)
	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// DeleteTask handles DELETE /tasks/{id} — deletes a task
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := h.service.DeleteTask(id)
	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, "Could not delete task", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent) // No content to return
}

// ListTasks handles GET /tasks — returns a list of tasks
func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	// Parse page and per_page (default: page=1, perPage=10)
	page, _ := strconv.Atoi(query.Get("page"))
	if page < 1 {
		page = 1
	}
	perPage, _ := strconv.Atoi(query.Get("per_page"))
	if perPage < 1 {
		perPage = 10
	}

	// Filter by status if provided
	filter := models.TaskFilter{
		Status: query.Get("status"),
	}

	tasks, err := h.service.ListTasks(page, perPage, filter)
	if err != nil {
		http.Error(w, "Could not fetch tasks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}