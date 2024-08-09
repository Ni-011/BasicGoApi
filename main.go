package main

import (
	"fmt"
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

	app.Patch("todos/:id", updateTodo)

	app.Delete("todos/:id", deleteTodo)

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

	todo.ID = len(todos) + 1     // update the id
	todos = append(todos, *todo) // append the new todo into todos

	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id") // get the id from the url

	for i, todo := range todos { // if id matches an id in todo, mark it as completed
		if fmt.Sprint(todo.ID) == id {
			todos[i].Completed = true
			return c.Status(200).JSON(todos[i])
		}
	}
	return c.Status(404).JSON(fiber.Map{"msg": "Todo not found"})
}

func deleteTodo (c *fiber.Ctx) error {
	id := c.Params("id")
	
	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.Status(200).JSON(fiber.Map{"msg": "Successfully deleted"})
		}
	}
	return c.Status(404).JSON(fiber.Map{"msg": "Todo not found"})
}
