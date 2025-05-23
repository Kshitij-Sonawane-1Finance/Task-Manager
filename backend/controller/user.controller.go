package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kshitij/taskManager/dto"
	"github.com/kshitij/taskManager/models"
	"github.com/kshitij/taskManager/services"
)

func CreateUser(ctx *fiber.Ctx) error {

	var user models.User;
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}
	
	result := services.CreateUser(ctx, user);

	return ctx.Status(result.StatusCode).JSON(result);

}


func Login(ctx *fiber.Ctx) error {

	var userLogin dto.UserLogin;

	err := ctx.BodyParser(&userLogin);
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}

	result := services.Login(ctx, userLogin);

	return ctx.Status(result.StatusCode).JSON(result);

}