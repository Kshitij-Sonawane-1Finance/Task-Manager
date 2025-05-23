package router

import "github.com/gofiber/fiber/v2"

func RouteHandler() *fiber.App {
	app := fiber.New()

	TaskRouter(app);
	UserRouter(app);

	return app;
}