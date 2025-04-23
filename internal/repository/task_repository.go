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

// Update modifies an existing task
func (r *InMemoryTaskRepository) Update(task *models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if task == nil {
		return errors.New("task is nil")
	}

	_, exists := r.tasks[task.ID]
	if !exists {
		return errors.New("task not found")
	}

	task.UpdatedAt = time.Now()
	r.tasks[task.ID] = *task

	return nil
}

// List returns a paginated and optionally filtered list of tasks
func (r *InMemoryTaskRepository) List(page, perPage int, filter models.TaskFilter) (*models.TaskPage, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var filteredTasks models.Tasks
	for _, task := range r.tasks {
		// Apply status filter if specified
		if filter.Status != "" && task.Status != filter.Status {
			continue
		}
		filteredTasks = append(filteredTasks, task)
	}

	totalCount := len(filteredTasks)

	// Handle invalid pagination inputs
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}

	start := (page - 1) * perPage
	end := start + perPage
	if start >= totalCount {
		return &models.TaskPage{
			Tasks:      models.Tasks{},
			TotalCount: totalCount,
			Page:       page,
			PerPage:    perPage,
		}, nil
	}
	if end > totalCount {
		end = totalCount
	}

	pagedTasks := filteredTasks[start:end]

	return &models.TaskPage{
		Tasks:      pagedTasks,
		TotalCount: totalCount,
		Page:       page,
		PerPage:    perPage,
	}, nil
}


// Delete removes a task by ID
func (r *InMemoryTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.tasks[id]
	if !exists {
		return errors.New("task not found")
	}

	delete(r.tasks, id)
	return nil
}
