# CRUD API - Task Manager for AI Agents

A simple CRUD API built with Go's standard library and MongoDB for managing AI agent tasks.

## Prerequisites

- Go 1.16 or higher
- MongoDB running locally on port 27017 (or use a connection string)

## Setup

1. Install the MongoDB driver:
```bash
go get go.mongodb.org/mongo-driver/mongo
```

2. Start MongoDB (if running locally):
```bash
# Using Docker
docker run -d -p 27017:27017 --name mongodb mongo:latest

# Or if installed locally
mongod
```

3. Run the application:
```bash
go run .
```

4. (Optional) Set custom MongoDB URI:
```bash
export MONGODB_URI="mongodb://localhost:27017"
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

Get a specific task (replace `<task_id>` with actual MongoDB ObjectID):
```bash
curl http://localhost:8080/tasks/<task_id>
```

Update a task (replace `<task_id>` with actual MongoDB ObjectID):
```bash
curl -X PUT http://localhost:8080/tasks/<task_id> \
  -H "Content-Type: application/json" \
  -d '{"title":"Analyze logs","description":"Parse error logs and generate report","status":"completed"}'
```

Delete a task (replace `<task_id>` with actual MongoDB ObjectID):
```bash
curl -X DELETE http://localhost:8080/tasks/<task_id>
```

## Database

The application uses MongoDB with:
- Database: `crud_api`
- Collection: `tasks`

## Environment Variables

- `MONGODB_URI`: MongoDB connection string (default: `mongodb://localhost:27017`)
