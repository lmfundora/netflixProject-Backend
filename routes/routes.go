package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lmfundora/netflixProject/handlers"

	"github.com/lmfundora/netflixProject/middlewares"
	"github.com/lmfundora/netflixProject/config"
)



func SetupRoutes(app *fiber.App) {
	jwt := middlewares.NewAuthMiddleware(config.Secret)


	app.Get("/", jwt, handlers.ListUsers)

	app.Get("/user/:id", jwt, handlers.ShowUser)

	app.Post("/user", handlers.CreateUsers)

	app.Patch("/user/:id", jwt, handlers.UpdateUser)

	app.Delete("/user/:id", jwt, handlers.DeleteUser)

	app.Post("/user/login", handlers.LogInUser)

}