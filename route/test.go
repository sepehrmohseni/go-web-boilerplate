package routes

import (
	"github.com/sepehrmohseni/go-web-boilerplate/controllers"

	"github.com/gofiber/fiber/v2"
)

func TestRoutes(testRoute fiber.Router) {
	// add biz info
	testRoute.Post("/", controllers.AddSomething)
}
