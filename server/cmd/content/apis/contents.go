package apis

import (
	"fmt"
	mysql "internal/databases"

	utils "github.com/zkfmapf123/go-js-utils"

	"github.com/gofiber/fiber/v2"
)

func GetPost() fiber.Handler{
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		p := mysql.NewPost()
		err := p.GetPost(id)

		if err != nil {
			return c.JSON(fiber.Map{
				"message" : fmt.Sprintf("Post not found : %s", err.Error()),
			})
		}

		return c.JSON(p)
	}
}

func PostCreate() fiber.Handler{
	return func(c *fiber.Ctx) error {
		body :=  utils.JsonParse[mysql.Post](c.Body())
		
		p := mysql.NewPost()
		p.MustGenerate(body.UserId, body.Title, body.Content, body.CategoryId)		

		return c.SendStatus(200)
	}
}

func PostUpdate() fiber.Handler{
	return func(c *fiber.Ctx) error {
		return c.SendString("Content")
	}
}

func PostDelete() fiber.Handler{
	return func(c *fiber.Ctx) error {
		return c.SendString("Content")
	}
}