package controllers

import (
	"fmt"
	"log"

	"github.com/Kamva/mgm/v2"
	"github.com/adimyth/go-fiber-crud/models"
	"github.com/adimyth/go-fiber-crud/schema"
	"github.com/adimyth/go-fiber-crud/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllTodos godoc
// @Summary      Get all todos
// @Description  Get all todos
// @Tags         crud todos
// @Accept       */*
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /todos [get]
func GetAllTodos(ctx *fiber.Ctx) error {
	// Access the collection
	collection := mgm.Coll(&models.ToDo{})
	// Fetch all todos & store it in todos variable
	todos := []models.ToDo{}
	err := collection.SimpleFind(&todos, bson.M{})
	// If there is an error, return the error
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "SERVER ERROR",
			"error":   err.Error(),
		})
	}
	// Otherwise, return the todos
	return ctx.JSON(fiber.Map{
		"ok":   true,
		"data": todos,
	})
}

// GetTodoByID godoc
// @Summary      Get todo by id
// @Description  Get todo by id
// @Tags         crud todos
// @Accept       */*
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /todos/:id [get]
func GetTodoByID(ctx *fiber.Ctx) error {
	collection := mgm.Coll(&models.ToDo{})
	todo := models.ToDo{}
	err := collection.FindByID(ctx.Params("id"), &todo)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"ok":      false,
			"message": "INVALID ID",
			"error":   err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"ok":   true,
		"data": todo,
	})
}

// CreateTodo godoc
// @Summary      Create a todo
// @Description  Create a todo. Provide title & description
// @Tags         crud todos
// @Accept       */*
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /todos [post]
func CreateTodo(ctx *fiber.Ctx) error {
	body := schema.ToDo{}

	// Parse the body. Return 400 if there is an error
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":      false,
			"message": "CANNOT PARSE BODY",
			"error":   err.Error(),
		})
	}

	// Validate request body
	errors := utils.ValidateStruct(body)
	if errors != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Get collection
	collection := mgm.Coll(&models.ToDo{})
	// Create todo from request parameters
	todo := models.CreateTodo(body.Title, body.Description)
	// Insert into collection
	err := collection.Create(todo)
	// Return 500 if there is an error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "SERVER ERRROR",
			"error":   err.Error(),
		})
	}
	// Return the created data otherwise
	return ctx.JSON(fiber.Map{
		"ok":   true,
		"data": body,
	})
}

// UpdateTodo godoc
// @Summary      Update a todo
// @Description  Update a todo. Provide title, description & id
// @Tags         crud todos
// @Accept       */*
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /todos/:id [patch]
func UpdateTodo(ctx *fiber.Ctx) error {
	// Extract id from params
	id := ctx.Params("id")
	body := schema.ToDo{}

	todoModel := models.ToDo{}
	// Get collection
	collection := mgm.Coll(&todoModel)
	// Get todo by id
	err := collection.FindByID(id, &todoModel)
	// If there is an error, return 404 (invalid id or not found)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"ok":      false,
			"error":   "INVALID ID",
			"message": err.Error(),
		})
	}

	// Parse the body. Return 400 if there is an error
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":      false,
			"error":   "CANNOT PARSE BODY",
			"message": err.Error(),
		})
	}

	// Validate request body
	errors := utils.ValidateStruct(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Update todo
	todoModel.Title = body.Title
	todoModel.Description = body.Description
	// Update the todo
	err = collection.Update(&todoModel)
	// Return 500 if there is an error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "SERVER ERROR",
			"error":   err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"ok":   true,
		"data": todoModel,
	})
}

// DeleteTodo godoc
// @Summary      Delete a todo
// @Description  Provide an id to delete a todo
// @Tags         crud todos
// @Accept       */*
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /todos/:id [delete]
func DeleteTodo(ctx *fiber.Ctx) error {
	// Extract id from params
	id := ctx.Params("id")

	todo := models.ToDo{}
	// Get collection
	collection := mgm.Coll(&todo)
	// Get todo by id
	err := collection.FindByID(id, &todo)
	// If there is an error, return 404 (invalid id or not found)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"ok":      false,
			"message": "INVALID ID",
			"error":   err.Error(),
		})
	}

	// Delete todo
	err = collection.Delete(&todo)
	// Return 500 if there is an error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"ok":      false,
			"message": "SERVER ERROR",
			"error":   err.Error(),
		})
	}
	msg := fmt.Sprintf("Todo (%v) deleted ", ctx.Params("id"))
	return ctx.SendString(msg)
}

// StatusVerification godoc
// @Summary      Middlewares Chaining demonstration
// @Description  Middleware chaining
// @Tags         verification
// @Accept       */*
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /verify/:status/:role [get]
func StatusVerification(c *fiber.Ctx) error {
	log.Println("Verifying user")
	if c.Locals("isAuthenticated") == false {
		return c.Status(403).SendString("Unauthenticated! Please sign up!")
	}
	return c.Status(302).SendString("Redirecting " + c.Locals("redirectRoute").(string))
}

// HealthCheck godoc
// @Summary      Show the status of server.
// @Description  Get the status of server.
// @Tags         healthcheck
// @Accept       */*
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       / [get]
func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("It's working!")
}
