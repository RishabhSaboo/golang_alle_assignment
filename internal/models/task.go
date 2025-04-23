package models

import "time"

// Task represents a single task in the system
type Task struct {
	ID          string    `json:"id"`          // Unique identifier
	Title       string    `json:"title"`       // Short title of the task
	Description string    `json:"description"` // Detailed description
	Status      string    `json:"status"`      // e.g., "Pending", "In Progress", "Completed"
	CreatedAt   time.Time `json:"created_at"`  // When the task was created
	UpdatedAt   time.Time `json:"updated_at"`  // When the task was last updated
}

// Tasks is a slice of Task
type Tasks []Task


type TaskFilter struct {
	Status string 
}

// TaskPage is a paginated response for tasks
type TaskPage struct {
	Tasks      Tasks `json:"tasks"`       
	TotalCount int   `json:"total_count"` 
	Page       int   `json:"page"`        
	PerPage    int   `json:"per_page"`    
}
