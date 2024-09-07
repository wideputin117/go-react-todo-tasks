package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// create a struct
type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("manush is")
	// this is how you create a server using fiber framework for go

	app := fiber.New()
	err := godotenv.Load(".env")
	// if err is not nil log a fatal error
	if err != nil {
		log.Fatal("The .env id not loades")
	}

	// GET PORT from the ENV
	PORT := os.Getenv("PORT")

	// route for psoting to todos
	// 1. create a todo variable to store

	todos := []Todo{}

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// attache the memory location of TOdo to todo
		todo := &Todo{}
		// parse it to c context which is the sent value by the user
		// also check for any error
		if error := c.BodyParser(todo); error != nil {
			return error
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todo.Completed = true
		todos = append(todos, *todo)
		return c.Status(200).JSON(todo)
	})

	log.Fatal(app.Listen(":" + PORT))
}
