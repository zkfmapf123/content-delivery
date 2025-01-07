package middlewares

import (
	"github.com/gofiber/fiber/v2"
	utils "github.com/zkfmapf123/go-js-utils"
)



func SchemaMiddleware[T any]() fiber.Handler {

	return func(c *fiber.Ctx) error {

		data := utils.JsonParse[T](c.Body())
		c.Locals("body", data)
		
		return c.Next()
	}
}