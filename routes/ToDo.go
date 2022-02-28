package routes

import (
	"github.com/adimyth/go-fiber-crud/controllers"
	_ "github.com/adimyth/go-fiber-crud/docs"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// CRUD operations
	api.Get("/todos", controllers.GetAllTodos)
	api.Get("/todos/:id", controllers.GetTodoByID)
	api.Post("/todos", controllers.CreateTodo)
	api.Patch("/todos/:id", controllers.UpdateTodo)
	api.Delete("/todos/:id", controllers.DeleteTodo)
}
