package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello, World!")
	app := fiber.New()

	// env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	// Slice to store todos
	todos := []Todo{}

	// Route to check if the server is working
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Route to get all todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Route to create a new todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := new(Todo)

		// Parse the request body into the Todo struct
		if err := c.BodyParser(todo); err != nil {
			return c.Status(422).JSON(fiber.Map{"msg": "Invalid request body", "error": err.Error()})
		}

		// Validate the body field
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"msg": "Todo body is required"})
		}

		// Assign an ID and add to the list
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		// Return the created todo
		return c.Status(201).JSON(todo)
	})

	// update a Todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"msg": "Todo not found"})
	})

	// Delete a todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"msg": "Todo deleted", "success": true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"msg": "Todo not found", "success": false})
	})

	log.Fatal(app.Listen(":" + PORT))
}
