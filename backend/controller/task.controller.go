package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kshitij/taskManager/dto"
	"github.com/kshitij/taskManager/models"
	"github.com/kshitij/taskManager/services"
)


func CreateTask(ctx *fiber.Ctx) error {

	var task models.Task

	err := ctx.BodyParser(&task)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}

	result := services.CreateTask(ctx, task)

	return ctx.Status(result.StatusCode).JSON(result);

}


func FindTask(ctx *fiber.Ctx) error {

	id := ctx.Params("id");

	idInt, err := strconv.ParseUint(id, 10, 64);
	if err != nil {
		log.Fatal(err);
	}

	result, err := services.FindTask(ctx, idInt);
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task does not Exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(result);

}

func FindAllTasks(ctx *fiber.Ctx) error {

	result, err := services.FindAllTasks(ctx);
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task does not Exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(result);

}

func DeleteTask(ctx *fiber.Ctx) error {

	id := ctx.Params("id");

	idInt, err := strconv.ParseUint(id, 10, 64);
	if err != nil {
		log.Fatal(err);
	}

	result := services.DeleteTask(ctx, idInt);

	return ctx.Status(result.StatusCode).JSON(result);
}

func SoftDelete(ctx *fiber.Ctx) error {

	id := ctx.Params("id");

	idInt, err := strconv.ParseUint(id, 10, 64);
	if err != nil {
		log.Fatal(err);
	}

	result := services.SoftDelete(ctx, idInt);

	return ctx.Status(result.StatusCode).JSON(result);
}


func UpdateTask(ctx *fiber.Ctx) error {

	id := ctx.Params("id");
	idInt, err := strconv.ParseUint(id, 10, 64);
	if err != nil {
		log.Fatal(err);
	}

	var updateTask dto.UpdateTask;

	err = ctx.BodyParser(&updateTask)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}

	result := services.UpdateTask(ctx, idInt, updateTask);

	return ctx.Status(result.StatusCode).JSON(result);

}