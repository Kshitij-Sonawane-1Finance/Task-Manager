package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kshitij/taskManager/controller"
)

func TaskRouter(app *fiber.App) {
	// app.Get("/", func (c *fiber.Ctx) error {
    //     return c.SendString("Hello, World!")
    // })

	taskApi := app.Group("/api/task")
	taskApi.Get("/", controller.FindAllTasks)
	taskApi.Post("/", controller.CreateTask)
	taskApi.Delete("/softDelete/:id", controller.SoftDelete)
	taskApi.Get("/:id", controller.FindTask)
	taskApi.Delete("/:id", controller.DeleteTask)
	taskApi.Put("/:id", controller.UpdateTask)
}