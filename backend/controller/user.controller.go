package controller

import (
	"github.com/gofiber/fiber/v2"
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


// Controllers
func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	
	return c.userService.CreateUser(ctx);

}


func (c *UserController) Login(ctx *fiber.Ctx) error {

	return c.userService.Login(ctx);

}


func (c *UserController) FetchUser(ctx *fiber.Ctx) error {

	userId := uint(ctx.Locals("user_id").(uint64))

	return c.userService.FetchUser(ctx, userId);

}