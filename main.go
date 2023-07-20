package main

import (
	"github.com/gofiber/fiber"
)

// create routes
func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Post(NewLeads)
	app.Delete(DeleteLeads)
}
func main() {
	app := fiber.New()

	setupRoutes(app)
}
