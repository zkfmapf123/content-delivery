package main

import (
	"fmt"
	"internal/net"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	SERVER_NAME = os.Getenv("SERVER_NAME")
	PORT = os.Getenv("PORT")
	VERSION = os.Getenv("VERSION")
)

func main() {
	app := net.NewHTTP(SERVER_NAME, VERSION)
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("OK %s", SERVER_NAME))
	})

	app.Post("/external", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "OK"})
	})

	net.Run(app, PORT)

}