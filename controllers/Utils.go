package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// StatusVerification godoc
// @Summary      Middlewares Chaining demonstration
// @Description  Middleware chaining
// @Tags         verification
// @Accept       json
// @Produce      json
// @Param		status path string true "authenticated / not authenticated"
// @Param		role path string true "admin / user / guest"
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
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       / [get]
func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("It's working!")
}

func DownloadFile(ctx *fiber.Ctx) error {
	return ctx.Download("./public/sample.pdf")
}
