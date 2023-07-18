package main

import (
	"github.com/Sanskarzz/go-k8s-ecs-restapi/database"
	"github.com/Sanskarzz/go-k8s-ecs-restapi/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("/facts", handlers.GetFacts)

	app.Post("/fact", handlers.CreateFact)

}
