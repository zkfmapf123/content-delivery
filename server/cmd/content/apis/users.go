package apis

import (
	"cmd/content/configs"
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

func PostUserCreate(broker string, topic string) fiber.Handler{
	return func(c *fiber.Ctx) error {

		body := utils.JsonParse[mysql.User](c.Body())
		
		u := mysql.NewUser()
		u.MustGenerate(0, body.Email, body.Password, false)


		// kafka
		producer, err := configs.NewKafkaProducer(broker, topic)
		if err != nil{
			return c.JSON(fiber.Map{
				"message" : fmt.Sprintf("Kafka Broker Error : %s", err.Error()),
			})
		}

		err = producer.SendMessage(u.Tuser)
		if err != nil {
			return c.JSON(fiber.Map{
				"message" : fmt.Sprintf("Kafka Producer Error : %s", err.Error()),
			})
		}

		defer producer.Close()

		return c.JSON(fiber.Map{
			"message" : "Create",
		})
	}
}

func PostUserDelete(broker string, topic string) fiber.Handler{
	return func(c *fiber.Ctx) error {
		_id := c.Params("id")
		id ,_ := strconv.Atoi(_id)

		u := mysql.NewUser()
		err := u.Delete(id)

		// Error
		if err != nil {
			return c.JSON(fiber.Map{
				"message" : fmt.Sprintf("User not found : %s", err.Error()),
			})
		}
		
		// Not Exists User
		if u.Tuser.Email == "" {
			return c.JSON(fiber.Map{
				"message" : fmt.Sprintf("User not found : %s", err.Error()),
			})
		}

		// kafka
		producer, err := configs.NewKafkaProducer(broker, topic)
		if err != nil{
			return c.JSON(fiber.Map{
				"message" : fmt.Sprintf("Kafka Broker Error : %s", err.Error()),
			})
		}

		err = producer.SendMessage(u.Tuser)
		if err != nil {
			return c.JSON(fiber.Map{
				"message" : fmt.Sprintf("Kafka Producer Error : %s", err.Error()),
			})
		}

		defer producer.Close()

		return c.JSON(fiber.Map{
			"message" : fmt.Sprintf("User delete user : %s", u.Tuser.Email),
		})
	}
}