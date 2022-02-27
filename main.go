package main

import (
	"log"

	"github.com/Kamva/mgm/v2"
	"github.com/adimyth/go-fiber-crud/controllers"
	_ "github.com/adimyth/go-fiber-crud/docs"
	"github.com/adimyth/go-fiber-crud/middlewares"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := mgm.SetDefaultConfig(nil, "todos", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error setting default config")
		log.Fatal(err)
	}
}

// @title           Fiber Swagger Example API
// @version         2.0
// @description     This is a sample server server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /
// @schemes   http
func main() {
	app := fiber.New()

	// HealthCheck
	app.Get("/", controllers.HealthCheck)

	// CRUD operations
	app.Get("/todos", controllers.GetAllTodos)
	app.Get("/todos/:id", controllers.GetTodoByID)
	app.Post("/todos", controllers.CreateTodo)
	app.Patch("/todos/:id", controllers.UpdateTodo)
	app.Delete("/todos/:id", controllers.DeleteTodo)

	// Authentication & Authorization Middlewares
	app.Get("/verify/:status/:role", middlewares.Authenticate, middlewares.Authorize, controllers.StatusVerification)

	// Static files
	app.Static("/static", "./public")

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":3000")
}
