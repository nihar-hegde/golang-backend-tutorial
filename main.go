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
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"mes":"Heelo world"})
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


	log.Fatal(app.Listen(":3000"))
}

