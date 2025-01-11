package main

import (
	"cmd/content/apis"
	"cmd/content/middlewares"
	"internal/net"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	SERVER_NAME = os.Getenv("SERVER_NAME")
	VERSION = os.Getenv("VERSION")	
	PORT = os.Getenv("PORT")
	BROKERS = os.Getenv("BROKERS")
	TOPIC = os.Getenv("TOPIC")
)

func main() {
	// Server
	app := net.NewHTTP(SERVER_NAME,VERSION)
	app.Use(logger.New())

	app.Get("/health",middlewares.LoggerMiddleware(), apis.HealthCheck())
	
	// User -> Kafka Producer
	app = apiUser(app)

	net.Run(app, PORT)	
}

func apiUser(app *fiber.App) *fiber.App{
	app.Get("/user/:id",middlewares.LoggerMiddleware(), apis.GetUser())
	app.Post("/user",middlewares.LoggerMiddleware(), apis.PostUserCreate(BROKERS,TOPIC))
	app.Delete("/user/:id",middlewares.LoggerMiddleware(), apis.PostUserDelete(BROKERS,TOPIC))
	return app
}

// func apiContent(app *fiber.App) *fiber.App{
// 	app.Get("/content/:id",middlewares.LoggerMiddleware(), apis.GetPost())
// 	app.Post("/content",middlewares.LoggerMiddleware(), apis.PostCreate())

// 	app.Post("/restore/:id",middlewares.LoggerMiddleware(), apis.PostRestore())
// 	app.Delete("/content/:id",middlewares.LoggerMiddleware(), apis.PostDelete())
// 	return app
// }