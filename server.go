package main

import (
	"log"
	// "os"

	"github.com/Kamva/mgm/v2"
	"github.com/adimyth/go-fiber-crud/controllers"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// mongoUrl := os.Getenv("MONGODB_URL")
	// mongoPort := os.Getenv("MONGODB_PORT")
	// mongoUser := os.Getenv("MONGODB_USER")
	// mongoPassword := os.Getenv("MONGODB_PASSWORD")
	// mongoDB := os.Getenv("MONGODB_DB")
	// err = mgm.SetDefaultConfig(nil, mongoDB, options.Client().ApplyURI(mongoUrl+":"+mongoPort).SetAuth(options.Credential{
	// 	Username: mongoUser,
	// 	Password: mongoPassword,
	// }))

	err = mgm.SetDefaultConfig(nil, "todos", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error setting default config")
		log.Fatal(err)
	}
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello World!")
	})

	app.Get("/todos", controllers.GetAllTodos)
	app.Get("/todos/:id", controllers.GetTodoByID)
	app.Post("/todos", controllers.CreateTodo)
	app.Patch("/todos/:id", controllers.UpdateTodo)
	app.Delete("/todos/:id", controllers.DeleteTodo)

	app.Listen(3000)
}
