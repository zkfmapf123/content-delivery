package apis

import (
	"fmt"
	mysql "internal/databases"
	"strconv"

	"github.com/gofiber/fiber/v2"
	utils "github.com/zkfmapf123/go-js-utils"
)

func GetUser() fiber.Handler{
	return func(c *fiber.Ctx) error {
		_id := c.Params("id")
		id, _ := strconv.Atoi(_id)
		
		u := mysql.NewUser()
		err := u.Get(id)
		
		if err != nil {
			return c.JSON(fiber.Map{
				"message" : fmt.Sprintf("User not found : %s", err.Error()),
			})
		}

		return c.JSON(u.Tuser)
	}
}

func PostUserCreate() fiber.Handler{
	return func(c *fiber.Ctx) error {

		body := utils.JsonParse[mysql.User](c.Body())
		
		u := mysql.NewUser()
		u.MustGenerate(0, body.Email, body.Password, false)

		return c.JSON(fiber.Map{
			"message" : "Create",
		})
	}
}

func PostUserDelete() fiber.Handler{
	return func(c *fiber.Ctx) error {
		_id := c.Params("id")
		id ,_ := strconv.Atoi(_id)

		u := mysql.NewUser()
		err := u.Delete(id)

		if err != nil {
			return c.JSON(fiber.Map{
				"message" : fmt.Sprintf("User not found : %s", err.Error()),
			})
		}

		return c.JSON(fiber.Map{
			"message" : "Hello World",
		})
	}
}