package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ToDo struct{
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main(){ 
	fmt.Println("Hello World")
	
	todos := []ToDo{}

	app := fiber.New();
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error{
		todo := &ToDo{} // default valuees will be set {ID:0, Completed:false, Body:""}

		if err := c.BodyParser(todo); err != nil{    // bodyParse will bind the request body to the ToDo struct
			return err;

		} 
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error":"Todo body is required"})
		}
		todo.ID = len(todos)+1;
		todos = append(todos, *todo)
		return c.Status(201).JSON(todo)

	})

	app.Patch("/api/todos/:id",func(c *fiber.Ctx) error{
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id{
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i]);
			}	
		}
		return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})
	})

	app.Delete("/api/todos/:id",func(c *fiber.Ctx) error{
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id{
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"message":"Todo deleted"})
			}	
		}
		return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})
	})


	log.Fatal(app.Listen(":3000"))
}

