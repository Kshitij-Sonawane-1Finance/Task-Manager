package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kshitij/taskManager/controller"
)

func TaskRouter(app *fiber.App) {
	// app.Get("/", func (c *fiber.Ctx) error {
    //     return c.SendString("Hello, World!")
    // })

	taskApi := app.Group("/api")
	taskApi.Post("/task", controller.CreateTask)
	taskApi.Get("/task/:id", controller.FindTask)
}