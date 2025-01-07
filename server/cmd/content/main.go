package main

import (
	"cmd/content/apis"
	"cmd/content/middlewares"
	"fmt"
	mysql "internal/databases"
	"internal/net"
	"os"

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

	// Config : mysql
	db, err := mysql.CreateDBConnection().
		WithHost(os.Getenv("DB_HOST")).
		WithPort(os.Getenv("DB_PORT")).
		WithUser(os.Getenv("DB_USER")).
		WithPassword(os.Getenv("DB_PASSWORD")).
		WithDatabase(os.Getenv("DB_DATABASE")).
		Build() 
	
	if err != nil {
		panic(err)
	}

	fmt.Println("db : ", db)

	app.Get("/health",middlewares.LoggerMiddleware(), apis.HealthCheck())
	net.Run(app, PORT)	
}
