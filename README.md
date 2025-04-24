# Task Management Microservice

This is a simple **Task Management System** built in Go. It allows users to create, read, update, and delete tasks. The system is designed as a microservice with clear separation of concerns.
The application uses in-memory storage for tasks but can be easily extended to use a database if needed. 
It supports pagination and filtering of tasks, allowing users to efficiently interact with task data.

## Setup

### Prerequisites
- [Go](https://golang.org/dl/)
- A terminal or command prompt

### Run the Service

```bash
git clone https://github.com/RishabhSaboo/rishabh_saboo_alle_assignment.git
cd rishabh_saboo_alle_assignment
go mod tidy
go build -o task-service cmd/server/main.go
./task-service
```

**The service will start on port 8081 by default.**



## API Documentation

### Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | `/tasks` | Create a new task |
| GET | `/tasks/{id}` | Get a task by ID |
| GET | `/tasks` | List tasks with pagination and filtering |
| PUT | `/tasks/{id}` | Update an existing task |
| DELETE | `/tasks/{id}` | Delete a task |

### Create a Task

**Request:**
```
POST /tasks
Content-Type: application/json

{
  "title": "Complete project documentation",
  "description": "Write comprehensive API documentation for the task service",
  "status": "Pending"
}
```

**Response:**
```
Status: 201 Created
Content-Type: application/json

{
  "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
  "title": "Complete project documentation",
  "description": "Write comprehensive API documentation for the task service",
  "status": "Pending",
  "created_at": "2023-04-18T14:30:45Z",
  "updated_at": "2023-04-18T14:30:45Z"
}
```

**Notes:**
- `title` is required
- If `status` is not provided, it defaults to "Pending"
- Valid status values are "Pending", "In Progress", and "Completed"

### Get a Task

**Request:**
```
GET /tasks/f47ac10b-58cc-4372-a567-0e02b2c3d479
```

**Response:**
```
Status: 200 OK
Content-Type: application/json

{
  "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
  "title": "Complete project documentation",
  "description": "Write comprehensive API documentation for the task service",
  "status": "Pending",
  "created_at": "2023-04-18T14:30:45Z",
  "updated_at": "2023-04-18T14:30:45Z"
}
```

### List Tasks

**Request:**
```
GET /tasks?page=1&per_page=10&status=Pending
```

**Parameters:**
- `page` (optional): Page number (default: 1)
- `per_page` (optional): Number of items per page (default: 10)
- `status` (optional): Filter by status ("Pending", "In Progress", "Completed")

**Response:**
```
Status: 200 OK
Content-Type: application/json

{
  "tasks": [
    {
      "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
      "title": "Complete project documentation",
      "description": "Write comprehensive API documentation for the task service",
      "status": "Pending",
      "created_at": "2023-04-18T14:30:45Z",
      "updated_at": "2023-04-18T14:30:45Z"
    }
  ],
  "total_count": 1,
  "page": 1,
  "per_page": 10
}
```

### Update a Task

**Request:**
```
PUT /tasks/f47ac10b-58cc-4372-a567-0e02b2c3d479
Content-Type: application/json

{
  "title": "Complete project documentation",
  "description": "Write comprehensive API documentation for the task service",
  "status": "In Progress"
}
```

**Response:**
```
Status: 200 OK
Content-Type: application/json

{
  "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
  "title": "Complete project documentation",
  "description": "Write comprehensive API documentation for the task service",
  "status": "In Progress",
  "created_at": "2023-04-18T14:30:45Z",
  "updated_at": "2023-04-18T14:35:12Z"
}
```

**Notes:**
- `title` is required
- Valid status values are "Pending", "In Progress", and "Completed"
- `created_at` is preserved from the original task

### Delete a Task

**Request:**
```
DELETE /tasks/f47ac10b-58cc-4372-a567-0e02b2c3d479
```

**Response:**
```
Status: 204 No Content
```

## Task Properties

| Field | Type | Description |
|-------|------|-------------|
| id | string | Unique identifier (UUID) |
| title | string | Title of the task (required) |
| description | string | Detailed description of the task |
| status | string | Current status (Pending, In Progress, Completed) |
| created_at | datetime | Creation timestamp |
| updated_at | datetime | Last update timestamp |

