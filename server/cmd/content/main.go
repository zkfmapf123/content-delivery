package main

import (
	"cmd/content/apis"
	"cmd/content/middlewares"
	"internal/net"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

const (
	PORT = "8080"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// Server
	app := net.NewHTTP("core", "1.0.0")
	app.Use(logger.New())
	
	app.Get("/health",middlewares.LoggerMiddleware(), apis.HealthCheck())
	net.Run(app, PORT)	
}
