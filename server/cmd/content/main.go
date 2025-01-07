package main

import (
	"cmd/content/apis"
	"cmd/content/middlewares"
	"internal/net"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	SERVER_NAME = "Content 발생서버"
	VERSION = "1.0.0"
	PORT = "8080"
)

func main() {
	// Server
	app := net.NewHTTP(SERVER_NAME,VERSION)
	app.Use(logger.New())

	app.Get("/health",middlewares.LoggerMiddleware(), apis.HealthCheck())
	
	app.Get("/content/:id",middlewares.LoggerMiddleware(), apis.GetPost())
	app.Post("/content",middlewares.LoggerMiddleware(), apis.PostCreate())
	app.Put("/content",middlewares.LoggerMiddleware(), apis.PostUpdate())
	app.Delete("/content",middlewares.LoggerMiddleware(), apis.PostDelete())

	net.Run(app, PORT)	
}
