package service

import (
	"errors"

	"github.com/pvnptl/task-service/internal/models"
	"github.com/pvnptl/task-service/internal/repository"
)

// TaskService defines the business logic for tasks
type TaskService interface {
	CreateTask(task *models.Task) error
	GetTaskByID(id string) (*models.Task, error)
	ListTasks(page, perPage int, filter models.TaskFilter) (*models.TaskPage, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id string) error
}


type DefaultTaskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a new service instance
func NewTaskService(repo repository.TaskRepository) *DefaultTaskService {
	return &DefaultTaskService{
		repo: repo,
	}
}

// CreateTask validates input and creates a new task
func (s *DefaultTaskService) CreateTask(task *models.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}

	if task.Title == "" {
		return errors.New("task title is required")
	}

	// Set default status if not provided
	if task.Status == "" {
		task.Status = "Pending"
	}


	if !isValidStatus(task.Status) {
		return errors.New("invalid task status")
	}

	return s.repo.Create(task)
}

// GetTaskByID fetches a task using its ID
func (s *DefaultTaskService) GetTaskByID(id string) (*models.Task, error) {
	return s.repo.GetByID(id)
}



// UpdateTask validates input and updates an existing task
func (s *DefaultTaskService) UpdateTask(task *models.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}

	if task.Title == "" {
		return errors.New("task title is required")
	}

	// Get existing task to preserve created time
	existingTask, err := s.repo.GetByID(task.ID)
	if err != nil {
		return err
	}
	if existingTask == nil {
		return errors.New("task not found")
	}

	if !isValidStatus(task.Status) {
		return errors.New("invalid task status")
	}

	task.CreatedAt = existingTask.CreatedAt

	return s.repo.Update(task)
}

// DeleteTask removes a task by ID
func (s *DefaultTaskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}

// ListTasks returns paginated and filtered task list
func (s *DefaultTaskService) ListTasks(page, perPage int, filter models.TaskFilter) (*models.TaskPage, error) {
	return s.repo.List(page, perPage, filter)
}


// isValidStatus checks if status is allowed
func isValidStatus(status string) bool {
	validStatuses := map[string]bool{
		"Pending":     true,
		"In Progress": true,
		"Completed":   true,
	}
	return validStatuses[status]
}