package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// InitDB initializes the MongoDB connection
func InitDB(uri string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database("crud_api").Collection("tasks")

	return &MongoDB{
		client:     client,
		collection: collection,
	}, nil
}

// Close closes the MongoDB connection
func (db *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.client.Disconnect(ctx)
}

// CreateTask inserts a new task into the database
func CreateTask(db *MongoDB, task *Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err := db.collection.InsertOne(ctx, task)
	return err
}

// GetAllTasks retrieves all tasks from the database
func GetAllTasks(db *MongoDB) ([]Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := db.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	// Return empty slice instead of nil if no tasks
	if tasks == nil {
		tasks = []Task{}
	}

	return tasks, nil
}

// GetTaskByID retrieves a single task by ID
func GetTaskByID(db *MongoDB, id string) (*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task Task
	err = db.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// UpdateTask updates an existing task
func UpdateTask(db *MongoDB, id string, task *Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	task.UpdatedAt = time.Now()
	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"status":      task.Status,
			"updated_at":  task.UpdatedAt,
		},
	}

	_, err = db.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	return err
}

// DeleteTask removes a task from the database
func DeleteTask(db *MongoDB, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = db.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
