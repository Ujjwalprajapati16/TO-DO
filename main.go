package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello, World!")
	app := fiber.New()

	// env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	MONGODB_URI := os.Getenv("MONGODB_URI")

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	collection = client.Database("todo_db").Collection("todos")

	// Route to check if the server is working
	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	// without database
	// // Slice to store todos
	// todos := []Todo{}

	// // Route to check if the server is working
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// // Route to get all todos
	// app.Get("/api/todos", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(todos)
	// })

	// // Route to create a new todo
	// app.Post("/api/todos", func(c *fiber.Ctx) error {
	// 	todo := new(Todo)

	// 	// Parse the request body into the Todo struct
	// 	if err := c.BodyParser(todo); err != nil {
	// 		return c.Status(422).JSON(fiber.Map{"msg": "Invalid request body", "error": err.Error()})
	// 	}

	// 	// Validate the body field
	// 	if todo.Body == "" {
	// 		return c.Status(400).JSON(fiber.Map{"msg": "Todo body is required"})
	// 	}

	// 	// Assign an ID and add to the list
	// 	todo.ID = len(todos) + 1
	// 	todos = append(todos, *todo)

	// 	// Return the created todo
	// 	return c.Status(201).JSON(todo)
	// })

	// // update a Todo
	// app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")

	// 	for i, todo := range todos {
	// 		if fmt.Sprint(todo.ID) == id {
	// 			todos[i].Completed = true
	// 			return c.Status(200).JSON(todos[i])
	// 		}
	// 	}
	// 	return c.Status(404).JSON(fiber.Map{"msg": "Todo not found"})
	// })

	// // Delete a todo
	// app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
	// 	id := c.Params("id")

	// 	for i, todo := range todos {
	// 		if fmt.Sprint(todo.ID) == id {
	// 			todos = append(todos[:i], todos[i+1:]...)
	// 			return c.Status(200).JSON(fiber.Map{"msg": "Todo deleted", "success": true})
	// 		}
	// 	}
	// 	return c.Status(404).JSON(fiber.Map{"msg": "Todo not found", "success": false})
	// })

	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Internal Server Error", "success": false})
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return c.Status(500).JSON(fiber.Map{"msg": "Internal Server Error", "success": false})
		}
		todos = append(todos, todo)
	}

	return c.Status(200).JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(422).JSON(fiber.Map{"msg": "Invalid request body", "error": err.Error()})
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"msg": "Todo body is required", "success": false})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Internal Server Error", "success": false})
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid ID", "success": false})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Internal Server Error", "success": false})
	}

	return c.Status(200).JSON(fiber.Map{"msg": "Todo updated", "success": true})
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "Invalid ID", "success": false})
	}

	filter := bson.M{"_id": objectID}

	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"msg": "Internal Server Error", "success": false})
	}

	return c.Status(200).JSON(fiber.Map{"msg": "Todo deleted", "success": true})
}
