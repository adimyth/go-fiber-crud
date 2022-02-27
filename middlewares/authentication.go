package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func Authenticate(c *fiber.Ctx) error {
	log.Println("Authenticating user")
	if c.Params("status") == "authenticated" {
		c.Locals("isAuthenticated", true)
	} else {
		c.Locals("isAuthenticated", false)
	}
	return c.Next()
}
