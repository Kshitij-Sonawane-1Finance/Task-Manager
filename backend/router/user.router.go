package router

import (
	"github.com/gofiber/fiber/v2"
	conn "github.com/kshitij/taskManager/connection"
	"github.com/kshitij/taskManager/controller"
	"github.com/kshitij/taskManager/loadEnv"
	"github.com/kshitij/taskManager/middleware"
	"github.com/kshitij/taskManager/services"
)

func UserRouter(app *fiber.App) {

	loadEnvService := loadEnv.NewLoadEnvService()

	// initialize the DB Service
	db := conn.NewDBService(loadEnvService);

	// Initialize the userService
	userService := services.NewUserService(db, loadEnvService);

	// Initialize the userController by passing in the userService as the dependency 
	userController := controller.NewUserController(userService);

	middlewareService := middleware.NewMiddlewareService(loadEnvService)

	userApi := app.Group("/api/user");
	userApi.Post("/register", userController.CreateUser);
	userApi.Post("/login", userController.Login);
	userApi.Get("/fetchUser", middlewareService.JWTMiddleWare, userController.FetchUser);

}