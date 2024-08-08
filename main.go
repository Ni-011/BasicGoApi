package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct { // Todo struct
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var todos []Todo = []Todo{} // storing Todos

func main() {
	app := fiber.New()

	app.Post("/todos", addTodo)

	log.Fatal(app.Listen(":8000"))
}

func addTodo(c *fiber.Ctx) error {
	todo := &Todo{}                            // a refference to a new Todo struct
	if err := c.BodyParser(todo); err != nil { // parse the body into the todo struct, if err return it
		return err
	}

	if todo.Body == "" { // todo must have a body
		return c.Status(400).JSON(fiber.Map{"msg": "Todo body cannot be empty"})
	}

	todo.ID = len(todos) + 1 // update the id
	todos = append(todos, *todo) // append the new todo into todos

	return c.Status(201).JSON(todo)
}
