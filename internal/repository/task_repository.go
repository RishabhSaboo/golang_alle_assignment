package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pvnptl/task-service/internal/models"
)

// TaskRepository defines methods for storing and retrieving tasks
type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id string) (*models.Task, error)
	List(page, perPage int, filter models.TaskFilter) (*models.TaskPage, error)
	Update(task *models.Task) error
	Delete(id string) error
}

// InMemoryTaskRepository is an in-memory implementation of TaskRepository
type InMemoryTaskRepository struct {
	tasks map[string]models.Task
	mu    sync.RWMutex
}

// NewInMemoryTaskRepository creates a new instance
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]models.Task),
	}
}

// Create adds a new task to the repository
func (r *InMemoryTaskRepository) Create(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if task == nil {
		return errors.New("task is nil")
	}

	// Generate a new UUID if not provided
	if task.ID == "" {
		task.ID = uuid.New().String()
	}

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	r.tasks[task.ID] = *task
	return nil
}

// GetByID retrieves a task by its ID
func (r *InMemoryTaskRepository) GetByID(id string) (*models.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

