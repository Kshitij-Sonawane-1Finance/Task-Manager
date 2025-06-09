package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RouteHandler() *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault));

	TaskRouter(app);
	UserRouter(app);

	return app;
}