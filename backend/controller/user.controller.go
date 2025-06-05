package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kshitij/taskManager/dto"
	"github.com/kshitij/taskManager/models"
	"github.com/kshitij/taskManager/services"
)

// We dont need an interface here, as there is no need for abstraction

// directly create a struct which will have the service as the dependency
type UserController struct {
	userService services.UserService
}

// constroller which expects a service to be passed in while calling, and it returns the controller methods
func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService};
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {

	var user models.User;
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}
	
	result := c.userService.CreateUser(ctx, user);

	return ctx.Status(result.StatusCode).JSON(result);

}


func (c *UserController) Login(ctx *fiber.Ctx) error {

	var userLogin dto.UserLogin;

	err := ctx.BodyParser(&userLogin);
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}

	result := c.userService.Login(ctx, userLogin);

	return ctx.Status(result.StatusCode).JSON(result);

}