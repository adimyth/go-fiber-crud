package controllers

import (
	"github.com/Kamva/mgm/v2"
	"github.com/adimyth/go-fiber-crud/models"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTodos(ctx *fiber.Ctx) {
	// Access the collection
	collection := mgm.Coll(&models.ToDo{})
	// Fetch all todos & store it in todos variable
	todos := []models.ToDo{}
	err := collection.SimpleFind(&todos, bson.M{})
	// If there is an error, return the error
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok": false,
			// "error": "Server error",
			"error": err.Error(),
		})
		// TODO: Check why is return important here
		return
	}
	// Otherwise, return the todos
	ctx.JSON(fiber.Map{
		"ok":   true,
		"data": todos,
	})
}

func GetTodoByID(ctx *fiber.Ctx) {
	collection := mgm.Coll(&models.ToDo{})
	todo := models.ToDo{}
	err := collection.FindByID(ctx.Params("id"), &todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid ID",
			// "error": err.Error(),
		})
		return
	}
	ctx.JSON(fiber.Map{
		"ok":   true,
		"data": todo,
	})
}

func CreateTodo(ctx *fiber.Ctx) {
	// Expected params structure
	params := new(struct {
		Title       string
		Description string
	})
	// Parse the body. Return 400 if there is an error
	if err := ctx.BodyParser(&params); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Cannot parse request",
			// "error": err.Error(),
		})
		return
	}

	// Return 400 if title is empty
	if params.Title == "" {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Title is required",
		})
		return
	} else if params.Description == "" {
		// Return 400 if description is empty
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Description is required",
		})
	}

	// Get collection
	collection := mgm.Coll(&models.ToDo{})
	// Create todo from request parameters
	todo := models.CreateTodo(params.Title, params.Description)
	// Insert into collection
	err := collection.Create(todo)
	// Return 500 if there is an error
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": "Server error",
			// "error": err.Error(),
		})
		return
	}
	// Return the created data otherwise
	ctx.JSON(fiber.Map{
		"ok":   true,
		"data": params,
	})
}

func UpdateTodo(ctx *fiber.Ctx) {
	// Extract id from params
	id := ctx.Params("id")

	todo := models.ToDo{}
	// Get collection
	collection := mgm.Coll(&todo)
	// Get todo by id
	err := collection.FindByID(id, &todo)
	// If there is an error, return 404 (invalid id or not found)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid ID",
			// "error": err.Error(),
		})
		return
	}

	// Expected params structure
	params := new(struct {
		Title       string
		Description string
	})
	// Parse the body. Return 400 if there is an error
	if err := ctx.BodyParser(&params); err != nil {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Cannot parse request",
			// "error": err.Error(),
		})
		return
	}

	// If title is empty, return 400
	if params.Title == "" {
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Title is required",
		})
		return
	} else if params.Description == "" {
		// If description is empty, return 400
		ctx.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Description is required",
		})
		return
	}

	// Update todo
	todo.Title = params.Title
	todo.Description = params.Description
	// Update the todo
	err = collection.Update(&todo)
	// Return 500 if there is an error
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": "Server error",
			// "error": err.Error(),
		})
		return
	}
	ctx.JSON(fiber.Map{
		"ok":   true,
		"data": params,
	})
}

func DeleteTodo(ctx *fiber.Ctx) {
	// Extract id from params
	id := ctx.Params("id")

	todo := models.ToDo{}
	// Get collection
	collection := mgm.Coll(&todo)
	// Get todo by id
	err := collection.FindByID(id, &todo)
	// If there is an error, return 404 (invalid id or not found)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid ID",
			// "error": err.Error(),
		})
		return
	}

	// Delete todo
	err = collection.Delete(&todo)
	// Return 500 if there is an error
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": "Server error",
			// "error": err.Error(),
		})
		return
	}
}
