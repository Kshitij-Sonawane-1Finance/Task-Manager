package controller

import (
	"fmt"
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

	task.UserID = uint(ctx.Locals("user_id").(uint64));
	result := services.CreateTask(ctx, task)

	return ctx.Status(result.StatusCode).JSON(result);

}


func FindTask(ctx *fiber.Ctx) error {

	id := ctx.Params("id");

	idInt, err := strconv.ParseUint(id, 10, 64);
	if err != nil {
		log.Fatal(err);
	}

	userID := uint(ctx.Locals("user_id").(uint64));
	result, err := services.FindTask(ctx, idInt, userID);
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task does not Exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(result);

}

func FindAllTasks(ctx *fiber.Ctx) error {

	userID := uint(ctx.Locals("user_id").(uint64));
	fmt.Println("User ID is: ", userID);

	result, err := services.FindAllTasks(ctx, userID);
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

	userID := uint(ctx.Locals("user_id").(uint64));
	result := services.DeleteTask(ctx, idInt, userID);

	return ctx.Status(result.StatusCode).JSON(result);
}

func SoftDelete(ctx *fiber.Ctx) error {

	id := ctx.Params("id");

	idInt, err := strconv.ParseUint(id, 10, 64);
	if err != nil {
		log.Fatal(err);
	}

	userID := uint(ctx.Locals("user_id").(uint64));
	result := services.SoftDelete(ctx, idInt, userID);

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

	userID := uint(ctx.Locals("user_id").(uint64));
	result := services.UpdateTask(ctx, idInt, updateTask, userID);

	return ctx.Status(result.StatusCode).JSON(result);

}