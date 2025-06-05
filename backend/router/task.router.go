package router

import (
	"github.com/gofiber/fiber/v2"
	conn "github.com/kshitij/taskManager/connection"
	"github.com/kshitij/taskManager/controller"
	"github.com/kshitij/taskManager/loadEnv"
	"github.com/kshitij/taskManager/middleware"
	"github.com/kshitij/taskManager/services"
)

func TaskRouter(app *fiber.App) {
	// app.Get("/", func (c *fiber.Ctx) error {
    //     return c.SendString("Hello, World!")
    // })

	loadEnvService := loadEnv.NewLoadEnvService()
	db := conn.NewDBService(loadEnvService);
	taskService := services.NewTaskService(db);
	taskController := controller.NewTaskController(taskService);
	middlewareService := middleware.NewMiddlewareService(loadEnvService);

	taskApi := app.Group("/api/task", middlewareService.JWTMiddleWare)
	taskApi.Get("/", taskController.FindAllTasks)
	taskApi.Post("/", taskController.CreateTask)
	taskApi.Delete("/softDelete/:id", taskController.SoftDelete)
	taskApi.Get("/:id", taskController.FindTask)
	taskApi.Delete("/:id", taskController.DeleteTask)
	taskApi.Put("/:id", taskController.UpdateTask)

}