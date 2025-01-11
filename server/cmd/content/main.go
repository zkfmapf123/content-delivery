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
)

func main() {
	// Server
	app := net.NewHTTP(SERVER_NAME,VERSION)
	app.Use(logger.New())

	app.Get("/health",middlewares.LoggerMiddleware(), apis.HealthCheck())
	
	// User
	app = apiUser(app)
	// Content
	app = apiContent(app)

	net.Run(app, PORT)	
}

func apiUser(app *fiber.App) *fiber.App{
	app.Get("/user/:id",middlewares.LoggerMiddleware(), apis.GetUser())
	app.Post("/user",middlewares.LoggerMiddleware(), apis.PostUserCreate())
	app.Delete("/user/:id",middlewares.LoggerMiddleware(), apis.PostUserDelete())
	return app
}

func apiContent(app *fiber.App) *fiber.App{
	app.Get("/content/:id",middlewares.LoggerMiddleware(), apis.GetPost())
	app.Post("/content",middlewares.LoggerMiddleware(), apis.PostCreate())

	app.Post("/restore/:id",middlewares.LoggerMiddleware(), apis.PostRestore())
	app.Delete("/content/:id",middlewares.LoggerMiddleware(), apis.PostDelete())
	return app
}