# Go Fiber CRUD App

## CRUD App

A simple todo crud app.

| Method | Endpoint     | Description    |
| ------ | ------------ | -------------- |
| GET    | `/api/todos`     | Get all todos  |
| GET    | `/api/todos/:id` | Get todo by id |
| POST   | `/api/todos`     | Create todo    |
| DELETE | `/api/todos/:id` | Delete todo    |
| PUT    | `/api/todos/:id` | Update todo    |
| GET    | `/` | HealthCheck    |
| GET    | `/static/golang.png` | Static File Access Example   |

1. `models/Todo.go` contains model definition
2. `controllers/Todo.go` contains the app logic

## Dependencies

1. [Go Fiber](https://gofiber.io) - An Express-inspired web framework written in Go.
2. [Mongo Models](https://github.com/Kamva/mgm) - Mongo Go Models (mgm) is a fast and simple MongoDB ODM for Go (based on official Mongo Go Driver)
3. [Air](https://github.com/cosmtrek/air) - Live reload for Go apps. Equivalent to `nodmeon` in Node.js
4. [Swagger Documentation](https://github.com/arsmn/fiber-swagger) - Swagger documentation for Go Fiber apps.

## Running the App

1. Clone the repository
2. Install required dependencies for the app & Air
   ```bash
   go mod tidy
   go get github.com/cosmtrek/air
   ```
3. Run the app
   ```bash
   air serve .
   ```

## Middleware

```go
// middlewares.go
package middlewares

func SimpleMiddleware(c *fiber.Ctx) error {
	log.Println("Simple Middleware")
	return c.Next()
}
```

```go
// controllers.go
package controllers

func SimpleController(c *fiber.Ctx) error {
	log.Println("Simple Controller")
	return c.Next()
}
```

```go
// main.go
package main

func main() {
	app := fiber.New()

	// Middleware -> Controller
	app.Get("/verify/:status/:role", middlewares.SimpleMiddleware,controllers.StatusVerification)

	app.Listen(":3000")
}
```

Refer middlewares directory for custom authentication & authorize middlewares.

## Schema Validation

Fiber supports schema validation for request body. Though it is nowhere as powerful / simple as Pydantic schema validation.

Refer `utils/Validators.go` for more details

```go
func CreateTodo(ctx *fiber.Ctx) error {
	body := schema.ToDo{}
	ctx.BodyParser(&body)

	// Validate request body
	errors := utils.ValidateStruct(body)
	if errors != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}
  ...
}

```

## Swagger Documentation

1. Install swag

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Generate swagger documentation. Find swag binary location.
   ```bash
   swag init main.go --output docs
   # /Users/adimyth/go/bin/swag init main.go --output docs
   ```
3. Format swagger documentation
   ```bash
   swag fmt
   ```
4. Download fiber-swagger
   ```bash
   go get github.com/arsmn/fiber-swagger
   ```
5. Add documentation for each endpoint.

   ```go
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
   ```

   Refer controllers for more examples

6. Add `swagger` endpoint in list of routes

```go
// main.go
package main

import (
  swagger "github.com/arsmn/fiber-swagger/v2"
  "github.com/gofiber/fiber/v2"
  // replace below with ur own docs
  _ "github.com/adimyth/go-fiber-crud/docs"
)

func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default
  ...
	app.Listen(":3000")
}
```

7. Open swagger documentation in browser - [swagger docs](http://localhost:3000/swagger)



## DB Setup

1. Install mongodb community version
2. Start mongodb
3. Create `todos` database & `todo` collection
4. Create a user with read & write access to `todos` database

   ```bash
   use todos;

   db.createCollection("todo");

   db.createUser({
     user: "test",
     pwd: "test123",
     roles: [{
       role: "readWrite",
       db: "todos"
     }]
   });
   ```

For some reason the authentication fails even when able to login via mongo shell. Have raised an [issue here](https://github.com/Kamva/mgm/issues/66)
