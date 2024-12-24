package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello, World!")
	app := fiber.New()

	// Slice to store todos
	todos := []Todo{}

	// Route to check if the server is working
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello, World!"})
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

	log.Fatal(app.Listen(":3000"))
}
