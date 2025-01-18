package middlewares

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware() fiber.Handler {

	return func(c *fiber.Ctx) error {
		log.Printf("Method : %s Path : %s", c.Method(),c.Path() )

		start := time.Now()

		logData := fiber.Map{
			"timestamp":     time.Now().Format(time.RFC3339),
			"status":       c.Response().StatusCode(),
			"method":       c.Method(),
			"path":         c.Path(),
			"latency_ms":   float64(time.Since(start).Milliseconds()),
			"ip":           c.IP(),
			"user_agent":   c.Get("User-Agent"),
			"request_id":   c.GetRespHeader("X-Request-Id"),
		}

		fmt.Printf("%+v\n", logData)
		
		err := c.Next()
		return err
	}
}