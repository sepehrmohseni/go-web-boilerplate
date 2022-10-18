package routes

import (
	"github.com/sepehrmohseni/go-web-boilerplate/controllers"

	"github.com/gofiber/fiber/v2"
)

func TestRoutes(testRoute fiber.Router) {
	// add something
	testRoute.Post("/", controllers.AddSomething)
}
