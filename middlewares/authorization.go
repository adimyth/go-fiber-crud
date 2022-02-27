package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func Authorize(c *fiber.Ctx) error {
	log.Println("Authorizing user")
	if c.Locals("isAuthenticated") == false {
		return c.Next()
	}

	if c.Params("role") == "admin" {
		c.Locals("redirectRoute", "admin")
	} else {
		c.Locals("redirectRoute", "contact-support")
	}
	return c.Next()
}
