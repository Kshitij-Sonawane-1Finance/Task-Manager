package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kshitij/taskManager/controller"
)

func UserRouter(app *fiber.App) {

	userApi := app.Group("/api/user");
	userApi.Post("/register", controller.CreateUser);
	userApi.Post("/login", controller.Login);

}