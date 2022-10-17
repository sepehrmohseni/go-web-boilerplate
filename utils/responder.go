package utils

import "github.com/gofiber/fiber/v2"

func ResponseHandler(c *fiber.Ctx, code int, success bool, msg string, data interface{}, pageNum uint16) error {
	return c.Status(code).JSON(map[string]interface{}{
		"success":     success,
		"msg":         msg,
		"data":        data,
		"page_number": pageNum,
	})
}
