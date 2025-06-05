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


type TaskController struct {
	taskService services.TaskService
}

func NewTaskController(taskService services.TaskService) *TaskController {
	return &TaskController{taskService}
}


func (c *TaskController) CreateTask(ctx *fiber.Ctx) error {

	var task models.Task
	err := ctx.BodyParser(&task)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}

	task.UserID = uint(ctx.Locals("user_id").(uint64));
	result := c.taskService.CreateTask(ctx, task)

	return ctx.Status(result.StatusCode).JSON(result);

}


func (c *TaskController) FindTask(ctx *fiber.Ctx) error {

	id := ctx.Params("id");

	idInt, err := strconv.ParseUint(id, 10, 64);
	if err != nil {
		log.Fatal(err);
	}

	userID := uint(ctx.Locals("user_id").(uint64));
	result, err := c.taskService.FindTask(ctx, idInt, userID);
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task does not Exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(result);

}

func (c *TaskController) FindAllTasks(ctx *fiber.Ctx) error {

	userID := uint(ctx.Locals("user_id").(uint64));
	fmt.Println("User ID is: ", userID);

	result, err := c.taskService.FindAllTasks(ctx, userID);
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Task does not Exist",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(result);

}

func (c *TaskController) DeleteTask(ctx *fiber.Ctx) error {

	id := ctx.Params("id");

	idInt, err := strconv.ParseUint(id, 10, 64);
	if err != nil {
		log.Fatal(err);
	}

	userID := uint(ctx.Locals("user_id").(uint64));
	result := c.taskService.DeleteTask(ctx, idInt, userID);

	return ctx.Status(result.StatusCode).JSON(result);
}

func (c *TaskController) SoftDelete(ctx *fiber.Ctx) error {

	id := ctx.Params("id");

	idInt, err := strconv.ParseUint(id, 10, 64);
	if err != nil {
		log.Fatal(err);
	}

	userID := uint(ctx.Locals("user_id").(uint64));
	result := c.taskService.SoftDelete(ctx, idInt, userID);

	return ctx.Status(result.StatusCode).JSON(result);
}


func (c *TaskController) UpdateTask(ctx *fiber.Ctx) error {

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
	result := c.taskService.UpdateTask(ctx, idInt, updateTask, userID);

	return ctx.Status(result.StatusCode).JSON(result);

}