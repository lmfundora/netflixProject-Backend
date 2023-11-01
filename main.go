package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	

    "github.com/gofiber/fiber/v2"
    "github.com/lmfundora/netflixProject/database"
    "github.com/lmfundora/netflixProject/routes"
)


func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
    	AllowOrigins: "https://gofiber.io, https://gofiber.net, http://localhost:3000",
    	AllowHeaders:  "Origin, Content-Type, Accept",
	}))


	routes.SetupRoutes(app)

	app.Listen(":3005")
}