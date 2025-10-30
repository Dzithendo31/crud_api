# CRUD API - Task Manager for AI Agents

A simple CRUD API built with Go's standard library and SQLite for managing AI agent tasks.

## Setup

1. Install the SQLite driver:
```bash
go get github.com/mattn/go-sqlite3
```

2. Run the application:
```bash
go run .
```

The server will start on `http://localhost:8080`

## API Endpoints

### Create a Task
```bash
POST /tasks
Content-Type: application/json

{
  "title": "Process user data",
  "description": "Extract and validate user information from CSV file",
  "status": "pending"
}
```

### Get All Tasks
```bash
GET /tasks
```

### Get a Single Task
```bash
GET /tasks/{id}
```

### Update a Task
```bash
PUT /tasks/{id}
Content-Type: application/json

{
  "title": "Process user data",
  "description": "Extract and validate user information from CSV file",
  "status": "completed"
}
```

### Delete a Task
```bash
DELETE /tasks/{id}
```

## Task Status Values
- `pending` - Task is waiting to be started
- `in_progress` - Task is currently being executed
- `completed` - Task has been finished

## Example Usage with curl

Create a task:
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Analyze logs","description":"Parse error logs and generate report","status":"pending"}'
```

Get all tasks:
```bash
curl http://localhost:8080/tasks
```

Get a specific task:
```bash
curl http://localhost:8080/tasks/1
```

Update a task:
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Analyze logs","description":"Parse error logs and generate report","status":"completed"}'
```

Delete a task:
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

## Database

The application uses SQLite with a file named `tasks.db` that will be created automatically when you run the application.
