package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sepehrmohseni/go-web-boilerplate/database"
	"github.com/sepehrmohseni/go-web-boilerplate/entities"
	"github.com/sepehrmohseni/go-web-boilerplate/utils"
)

func AddSomething(c *fiber.Ctx) error {
	tst := new(entities.Test)
	if err := c.BodyParser(&tst); err != nil {
		return utils.ResponseHandler(
			c,
			500,
			false,
			err.Error(),
			nil,
			0,
		)
	}
	createTest := database.Database.Create(&tst)
	if createTest.Error != nil {
		return utils.ResponseHandler(
			c,
			500,
			false,
			createTest.Error.Error(),
			nil,
			0,
		)
	}
	return utils.ResponseHandler(
		c,
		201,
		true,
		"Your Test added successfuly",
		&tst,
		0,
	)
}